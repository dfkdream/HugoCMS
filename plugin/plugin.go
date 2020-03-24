package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	reservedAddr          = "/live,/metadata"
	reservedAdminAPIAddr  = "/post,/list,/blob,/whoami,/build,/config"
	reservedAdminPageAddr = "/,/assets,/list,/edit,/config"
)

type contextKey string

var (
	// ContextKeyUser is context key for user data
	ContextKeyUser = contextKey("user")
)

var (
	// ReservedAddrConflictError is returned when assigning address which conflicts with reserved addresses.
	ReservedAddrConflictError = errors.New("address conflicts with reserved address")
)

// User contains user information
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// String converts user to json string
func (u User) String() string {
	if res, err := json.Marshal(u); err == nil {
		return string(res)
	} else {
		return ""
	}
}

// Info contains information about plugin which will be displayed on HugoCMS dashboard.
type Info struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type adminEndpoint struct {
	MenuName string `json:"menuName"`
	Endpoint string `json:"endpoint"`
}

// Metadata contains metadata about plugin
type Metadata struct {
	Info              Info            `json:"info"`
	AdminEndpoints    []adminEndpoint `json:"adminEndpoints"`
	AdminAPIEndpoints []string        `json:"adminAPIEndpoints"`
	APIEndpoints      []string        `json:"apiEndpoints"`
}

// Plugin is HugoCMS Plugin which implements http.Handler.
type Plugin struct {
	router   *mux.Router
	metadata *Metadata
}

// New creates new plugin.
func New(Info Info) *Plugin {
	p := &Plugin{
		router: mux.NewRouter().StrictSlash(true),
		metadata: &Metadata{
			Info:              Info,
			AdminEndpoints:    make([]adminEndpoint, 0),
			AdminAPIEndpoints: make([]string, 0),
			APIEndpoints:      make([]string, 0),
		},
	}

	p.router.HandleFunc("/metadata", func(res http.ResponseWriter, req *http.Request) {
		err := json.NewEncoder(res).Encode(p.metadata)
		if err != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	})

	p.router.HandleFunc("/live", func(res http.ResponseWriter, req *http.Request) {
	})
	return p
}

// HandleAdminPage handles admin page handlers.
// menuName will be displayed on navigation bar.
// Handler should write HTML document.
func (p *Plugin) HandleAdminPage(path, menuName string, handler http.Handler) {
	if strings.Contains(reservedAddr, path) || strings.Contains(reservedAdminPageAddr, path) {
		panic(ReservedAddrConflictError)
	}
	p.metadata.AdminEndpoints = append(p.metadata.AdminEndpoints, adminEndpoint{Endpoint: path, MenuName: menuName})
	p.router.Handle(path, handler)
}

// HandleAdminAPI handles admin only API handlers.
// Non logged in requests will be rejected.
func (p *Plugin) HandleAdminAPI(path string, handler http.Handler) {
	if strings.Contains(reservedAddr, path) || strings.Contains(reservedAdminAPIAddr, path) {
		panic(ReservedAddrConflictError)
	}
	p.metadata.AdminAPIEndpoints = append(p.metadata.AdminAPIEndpoints, path)
	p.router.Handle(path, handler)
}

// HandleAPI handles API handlers.
// Non logged in users can access these APIs.
func (p *Plugin) HandleAPI(path string, handler http.Handler) {
	if strings.Contains(reservedAddr, path) {
		panic(ReservedAddrConflictError)
	}
	p.metadata.APIEndpoints = append(p.metadata.APIEndpoints, path)
	p.router.Handle(path, handler)
}

// ServeHTTP dispatches the requests to plugin.
func (p *Plugin) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if h := req.Header.Get("X-HugoCMS-User"); h != "" {
		u := new(User)
		err := json.Unmarshal([]byte(h), &u)
		if err != nil {
			http.Error(res, "Bad Request", http.StatusBadRequest)
			log.Println(err)
			return
		}
		req = req.WithContext(context.WithValue(req.Context(), ContextKeyUser, u))
	}
	p.router.ServeHTTP(res, req)
}

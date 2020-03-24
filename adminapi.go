package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dfkdream/hugocms/plugin"

	"github.com/gorilla/mux"
)

type adminAPI struct {
	conf *config
	hugo *hugo
}

func (a adminAPI) postAPI(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		f, err := os.Open(filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusNotFound, http.StatusNotFound)
			return
		}
		defer func() { _ = f.Close() }()
		a, err := parseArticle(f)
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusNotFound, http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(res).Encode(a)
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
			return
		}
	case "POST":
		var articleJSON article
		err := json.NewDecoder(req.Body).Decode(&articleJSON)
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusBadRequest, http.StatusBadRequest)
			return
		}
		f, err := os.OpenFile(filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(0644))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
			return
		}
		defer func() { _ = f.Close() }()
		jsonEnc := json.NewEncoder(f)
		jsonEnc.SetIndent("", "    ")
		err = jsonEnc.Encode(articleJSON.FrontMatter)
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
			return
		}
		_, err = f.Write([]byte(articleJSON.Body))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
			return
		}
	default:
		http.Error(res, jsonStatusMethodNotAllowed, http.StatusMethodNotAllowed)
	}
}

type fileInfo struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	Mode    string    `json:"mode"`
	ModTime time.Time `json:"modTime"`
	IsDir   bool      `json:"isDir"`
}

func (a adminAPI) listAPI(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		files, err := ioutil.ReadDir(filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusNotFound, http.StatusNotFound)
			return
		}
		fJSON := make([]fileInfo, len(files))

		for idx, f := range files {
			fJSON[idx].Name = f.Name()
			fJSON[idx].Size = f.Size()
			fJSON[idx].Mode = f.Mode().String()
			fJSON[idx].ModTime = f.ModTime()
			fJSON[idx].IsDir = f.IsDir()
		}

		err = json.NewEncoder(res).Encode(fJSON)
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
			return
		}
	case "POST":
		err := os.MkdirAll(filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)), os.FileMode(0755))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
			return
		}
	case "PUT":
		var path string
		err := json.NewDecoder(req.Body).Decode(&path)
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusBadRequest, http.StatusBadRequest)
			return
		}
		err = os.Rename(filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)), filepath.Join(a.conf.ContentPath, filepath.Clean(path)))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
		}
	case "DELETE":
		err := os.RemoveAll(filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
		}
	default:
		http.Error(res, jsonStatusMethodNotAllowed, http.StatusMethodNotAllowed)
	}
}

func (a adminAPI) blobAPI(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if strings.HasSuffix("/", req.URL.Path) {
			http.Error(res, "404 page not found", http.StatusNotFound)
			return
		}
		res.Header().Del("Content-Type")
		http.ServeFile(res, req, filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)))
	case "POST":
		f, err := os.OpenFile(filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.FileMode(0644))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
			return
		}
		defer func() { _ = f.Close() }()

		_, err = io.Copy(f, req.Body)
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
		}
	case "PUT":
		var path string
		err := json.NewDecoder(req.Body).Decode(&path)
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusBadRequest, http.StatusBadRequest)
			return
		}
		err = os.Rename(filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)), filepath.Join(a.conf.ContentPath, filepath.Clean(path)))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
		}
	case "DELETE":
		err := os.Remove(filepath.Join(a.conf.ContentPath, filepath.Clean("/"+req.URL.Path)))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
		}
	default:
		http.Error(res, jsonStatusMethodNotAllowed, http.StatusMethodNotAllowed)
	}
}

func (a adminAPI) whoamiAPI(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if u, ok := req.Context().Value(contextKeyUser).(*user); ok {
			err := json.NewEncoder(res).Encode(
				struct {
					ID       string `json:"id"`
					Username string `json:"username"`
				}{
					ID:       u.ID,
					Username: u.Username,
				})
			if err != nil {
				log.Println(err)
				http.Error(res, jsonStatusForbidden, http.StatusForbidden)
			}
			return
		}
		http.Error(res, jsonStatusForbidden, http.StatusForbidden)
	default:
		http.Error(res, jsonStatusMethodNotAllowed, http.StatusMethodNotAllowed)
	}
}

func (a adminAPI) buildAPI(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		r := a.hugo.build()
		if r.err != nil {
			log.Println(r.err)
		}
		fmt.Println(r.Result)
		err := json.NewEncoder(res).Encode(r)
		if err != nil {
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
		}
	default:
		http.Error(res, jsonStatusMethodNotAllowed, http.StatusMethodNotAllowed)
	}
}

func (a adminAPI) configAPI(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		res.Header().Del("Content-Type")
		http.ServeFile(res, req, a.conf.ConfigPath)
	case "POST":
		f, err := os.OpenFile(a.conf.ConfigPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(0644))
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
			return
		}
		defer func() { _ = f.Close() }()
		_, err = io.Copy(f, req.Body)
		if err != nil {
			log.Println(err)
			http.Error(res, jsonStatusInternalServerError, http.StatusInternalServerError)
			return
		}
	default:
		http.Error(res, jsonStatusMethodNotAllowed, http.StatusMethodNotAllowed)
	}
}

func (a adminAPI) setupHandlers(router *mux.Router) {
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.Header().Set("Content-Type", "application/json")
			res.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			next.ServeHTTP(res, req)
		})
	})

	router.PathPrefix("/post").Handler(http.StripPrefix("/admin/api/post", http.HandlerFunc(a.postAPI)))
	router.PathPrefix("/list").Handler(http.StripPrefix("/admin/api/list", http.HandlerFunc(a.listAPI)))
	router.PathPrefix("/blob").Handler(http.StripPrefix("/admin/api/blob", http.HandlerFunc(a.blobAPI)))
	router.HandleFunc("/whoami", a.whoamiAPI)
	router.HandleFunc("/build", a.buildAPI)
	router.HandleFunc("/config", a.configAPI)

	for _, v := range a.conf.Plugins {
		for _, path := range v.Metadata.AdminAPIEndpoints {
			router.Handle(path, &httputil.ReverseProxy{Director: func(req *http.Request) {
				target, err := url.Parse(singleJoiningSlash(v.Addr, path))
				if err != nil {
					log.Fatal(err)
				}
				req.URL.Scheme = target.Scheme
				req.URL.Host = target.Host
				req.URL.Path = target.Path

				if target.RawQuery == "" || req.URL.RawPath == "" {
					req.URL.RawQuery = target.RawQuery + req.URL.RawQuery
				} else {
					req.URL.RawQuery = target.RawQuery + "&" + req.URL.RawQuery
				}

				if _, ok := req.Header["User-Agent"]; !ok {
					req.Header.Set("User-Agent", "")
				}

				req.Header.Del("X-HugoCMS-User")

				if u, ok := req.Context().Value(contextKeyUser).(*user); ok {
					req.Header.Set("X-HugoCMS-User", plugin.User{ID: u.ID, Username: u.Username}.String())
				}
			}})
		}
	}
}

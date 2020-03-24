package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"

	"github.com/dfkdream/hugocms/plugin"
)

func newAuthenticatedReverseProxy(addr, path string) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{Director: func(req *http.Request) {
		target, err := url.Parse(singleJoiningSlash(addr, path))
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
	}}
}

type pluginAPI struct {
	config *config
	signIn *signInHandler
}

func (p pluginAPI) setupHandlers(router *mux.Router) {
	router.Use(p.signIn.middleware(false))
	for _, v := range p.config.Plugins {
		for _, path := range v.Metadata.APIEndpoints {
			router.Handle(path, newAuthenticatedReverseProxy(v.Addr, path))
		}
	}
}

package duan

import (
	"log"
	"net/http"
	"regexp"
)

type MiddlewareFunc func(handler http.Handler) http.Handler

type Router struct {
	middlewareChain []MiddlewareFunc
	mux             map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		middlewareChain: []MiddlewareFunc{},
		mux:             make(map[string]http.Handler),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handled := false
	for route, handler := range r.mux {
		matched, err := regexp.MatchString(route, req.URL.Path)
		if err != nil {
			http.Error(w, "route error", http.StatusInternalServerError)
			return
		}

		if matched {
			handled = true
			handler.ServeHTTP(w, req)
			break
		}
	}

	if !handled {
		log.Println("ERROR: no handler find: ", req.URL.Path)
	}
}

func (r *Router) Use(m MiddlewareFunc) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Path(route string, handler http.Handler) {
	var mergedHandler = handler
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}
	r.mux[route] = mergedHandler
}

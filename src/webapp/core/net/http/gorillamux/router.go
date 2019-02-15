package gorillamux

import (
	"net/http"
	wrapper "webapp/core/net/http"

	"github.com/gorilla/mux"
)

// Router to handle http requests
type Router struct {
	router *mux.Router
}

// New Creates new Gorilla mux router
func New() *Router {
	return &Router{
		router: mux.NewRouter(),
	}
}

// Handler return a handler created
func (r *Router) Handler() http.Handler {
	return r.router
}

// HandleRoute set the router
func (r *Router) HandleRoute(routes ...wrapper.Route) {
	for _, route := range routes {
		r.router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}

//Static add static context to the router
func (r *Router) Static(path string, root string) {
	r.router.PathPrefix(path).
		Handler(http.StripPrefix(path, http.FileServer(http.Dir(root))))
}

// Use set the middleware to use by default
func (r *Router) Use(mw ...wrapper.Middleware) {
	for _, m := range mw {
		r.router.Use(mux.MiddlewareFunc(m))
	}
}

//Vars get vars from a request
func (r *Router) Vars(req *http.Request) map[string]string {
	return mux.Vars(req)
}

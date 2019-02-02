package router

import (
	"fmt"
	"net/http"
	wrapper "webapp/core/net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Router to handle http requests
type Router struct {
	router     *httprouter.Router
	middleware alice.Chain
	routes     map[string]wrapper.Route
}

// ContextHandler to wrapper with gorilla mux
func ContextHandler(h http.Handler) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, p httprouter.Params) {

		h.ServeHTTP(res, req)
	}
}

// New Creates new Gorilla mux router
func New() *Router {
	return &Router{
		router:     httprouter.New(),
		middleware: alice.New(),
		routes:     map[string]wrapper.Route{},
	}
}

// Handler return a handler created
func (r *Router) Handler() http.Handler {
	return r.router
}

// HandleRoute set the router
func (r *Router) setRoutes(routes ...wrapper.Route) {
	for _, route := range routes {
		r.router.Handler(route.Method, route.Path, r.middleware.ThenFunc(http.HandlerFunc(route.Handler)))
	}
}

// HandleRoute set the router
func (r *Router) routeID(route wrapper.Route) string {
	return fmt.Sprintf("%s:%s", route.Method, route.Path)
}

// HandleRoute set the router
func (r *Router) HandleRoute(routes ...wrapper.Route) {
	for _, route := range routes {
		r.routes[r.routeID(route)] = route
	}
	r.setRoutes(routes...)
}

// Use set the middleware to use by default
func (r *Router) Use(mw ...wrapper.Middleware) {
	for _, m := range mw {
		r.middleware.Append(alice.Constructor(m))
	}
	//r.setRoutes()
}

//Vars get vars from a request
func (r *Router) Vars(req *http.Request) map[string]string {
	//return mux.Vars(req)
	return nil
}

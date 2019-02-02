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
	middleware []alice.Constructor
}

// New Creates new Gorilla mux router
func New() *Router {
	return &Router{
		router:     httprouter.New(),
		middleware: []alice.Constructor{},
	}
}

// Handler return a handler created
func (r *Router) Handler() http.Handler {
	return alice.New(r.middleware...).Then(r.router)
}

// HandleRoute set the router
func (r *Router) routeID(route wrapper.Route) string {
	return fmt.Sprintf("%s:%s", route.Method, route.Path)
}

// HandleRoute set the router
func (r *Router) HandleRoute(routes ...wrapper.Route) {
	for _, route := range routes {
		r.router.Handler(route.Method, route.Path, http.HandlerFunc(route.Handler))
	}
}

// Use set the middleware to use by default
func (r *Router) Use(mw ...wrapper.Middleware) {
	for _, m := range mw {
		r.middleware = append(r.middleware, alice.Constructor(m))
	}
}

//Vars get vars from a request
func (r *Router) Vars(req *http.Request) map[string]string {
	//return mux.Vars(req)
	return nil
}

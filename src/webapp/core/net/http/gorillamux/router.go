package gorillamux

import (
	"context"
	"net/http"
	wrapper "webapp/core/net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Router to handle http requests
type Router struct {
	router     *mux.Router
	middleware []wrapper.Middleware
	routes     []wrapper.Route
}

// New Creates new Gorilla mux router
func New() *Router {
	return &Router{
		router:     mux.NewRouter(),
		middleware: make([]wrapper.Middleware, 0),
		routes:     make([]wrapper.Route, 0),
	}
}

// Handler return a handler created
func (r *Router) Handler() http.Handler {
	global, filters := wrapper.SplitMiddleware(r.middleware)
	// Create chained handler for global middleware
	for _, m := range global {
		r.router.Use(mux.MiddlewareFunc(m.Handler()))
	}
	// Create chained handler for filter middleware
	fm := make([]alice.Constructor, 0, len(filters))
	for _, m := range filters {
		fm = append(fm, alice.Constructor(m.Handler()))
	}
	// Per route append the chained filter, with route information
	for _, route := range r.routes {
		r.router.HandleFunc(route.Path,
			wrapHandler(alice.New(fm...).
				Then(http.HandlerFunc(route.Handler)), route)).
			Methods(route.Method)
	}
	return r.router
}

// HandleRoute set the router
func (r *Router) HandleRoute(routes ...wrapper.Route) {
	r.routes = append(r.routes, routes...)
}

//Static add static context to the router
func (r *Router) Static(path string, root string) {
	r.router.PathPrefix(path).
		Handler(http.StripPrefix(path, http.FileServer(http.Dir(root))))
}

// Use set the middleware to use by default
func (r *Router) Use(mw ...wrapper.Middleware) {
	r.middleware = append(r.middleware, mw...)

}

//Vars get vars from a request
func (r *Router) Vars(req *http.Request) map[string]string {
	return mux.Vars(req)
}

func wrapHandler(h http.Handler, route wrapper.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, wrapper.RouteInfoKey, route)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

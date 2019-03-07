package gorillamux

import (
	"context"
	"net/http"
	net "webapp/core/net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Router to handle http requests
type Router struct {
	router     *mux.Router
	middleware []net.Middleware
	routes     []net.Route
}

// New Creates new Gorilla mux router
func New() *Router {
	return &Router{
		router:     mux.NewRouter(),
		middleware: make([]net.Middleware, 0),
		routes:     make([]net.Route, 0),
	}
}

// Handler return a handler created
func (r *Router) Handler() http.Handler {
	global, filters := net.SplitMiddleware(r.middleware)
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
func (r *Router) HandleRoute(routes ...net.Route) {
	r.routes = append(r.routes, routes...)
}

//Static add static context to the router
func (r *Router) Static(path string, root string) {
	r.router.PathPrefix(path).
		Handler(http.StripPrefix(path, http.FileServer(http.Dir(root))))
}

// Use set the middleware to use by default
func (r *Router) Use(mw ...net.Middleware) {
	r.middleware = append(r.middleware, mw...)

}

//Vars get vars from a request
func (r *Router) Vars(req *http.Request) map[string]string {
	return mux.Vars(req)
}

func wrapHandler(h http.Handler, route net.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, net.RouteInfoKey, route)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

package httprouter

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	log "webapp/core/logging"
	wrapper "webapp/core/net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// ContextTypeKey type used to add params infor to the reques context
type ContextTypeKey string

const paramskey ContextTypeKey = "params"

// Router to handle http requests
type Router struct {
	router     *httprouter.Router
	middleware []wrapper.Middleware
	routes     []wrapper.Route
}

// New Creates new Gorilla mux router
func New() *Router {
	return &Router{
		router:     httprouter.New(),
		middleware: make([]wrapper.Middleware, 0),
		routes:     make([]wrapper.Route, 0),
	}
}

// Handler return a handler created
func (r *Router) Handler() http.Handler {
	global, filters := wrapper.SplitMiddleware(r.middleware)
	// Create chained handler for global middleware
	gm := make([]alice.Constructor, 0, len(global))
	for _, m := range global {
		gm = append(gm, alice.Constructor(m.Handler()))
	}
	// Create chained handler for filter middleware
	fm := make([]alice.Constructor, 0, len(filters))
	for _, m := range filters {
		fm = append(fm, alice.Constructor(m.Handler()))
	}
	// Per route append the chained filter, with route information
	for _, route := range r.routes {
		r.router.Handle(route.Method, r.normalize(route.Path),
			wrapHandler(alice.New(fm...).Then(http.HandlerFunc(route.Handler)), route))
	}
	return alice.New(gm...).Then(r.router)
}

// HandleRoute set the router
func (r *Router) routeID(route wrapper.Route) string {
	return fmt.Sprintf("%s:%s", route.Method, route.Path)
}

func (r *Router) normalize(path string) string {
	return strings.Replace(strings.Replace(path, "}", "", -1), "{", ":", -1)
}

// HandleRoute set the router
func (r *Router) HandleRoute(routes ...wrapper.Route) {
	r.routes = append(r.routes, routes...)
}

//Static add static context to the router
func (r *Router) Static(path string, root string) {
	r.router.ServeFiles(path+"*filepath", http.Dir(root))
}

// Use set the middleware to use by default
func (r *Router) Use(mw ...wrapper.Middleware) {
	r.middleware = append(r.middleware, mw...)
}

//Vars get vars from a request
func (r *Router) Vars(req *http.Request) map[string]string {
	result := map[string]string{}
	ps, ok := req.Context().Value(paramskey).(httprouter.Params)
	if !ok {
		return result
	}
	for _, item := range ps {
		result[item.Key] = item.Value
	}
	log.Debug(result)
	return result
}

func wrapHandler(h http.Handler, route wrapper.Route) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, paramskey, ps)
		ctx = context.WithValue(ctx, wrapper.RouteInfoKey, route)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

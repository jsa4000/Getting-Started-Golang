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

const params = "params"

// Router to handle http requests
type Router struct {
	router     *httprouter.Router
	middleware []wrapper.Middleware
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, params, ps)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}

// New Creates new Gorilla mux router
func New() *Router {
	return &Router{
		router:     httprouter.New(),
		middleware: []wrapper.Middleware{},
	}
}

// Handler return a handler created
func (r *Router) Handler() http.Handler {
	sm := wrapper.SortMiddleware(r.middleware, true)
	am := make([]alice.Constructor, 0, len(sm))
	for _, m := range sm {
		am = append(am, alice.Constructor(m.Handler()))
	}
	return alice.New(am...).Then(r.router)
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
	for _, route := range routes {
		route.Path = r.normalize(route.Path)
		switch strings.ToLower(route.Method) {
		case "get":
			r.router.GET(route.Path, wrapHandler(http.HandlerFunc(route.Handler)))
		case "post":
			r.router.POST(route.Path, wrapHandler(http.HandlerFunc(route.Handler)))
		case "put":
			r.router.PUT(route.Path, wrapHandler(http.HandlerFunc(route.Handler)))
		case "patch":
			r.router.PATCH(route.Path, wrapHandler(http.HandlerFunc(route.Handler)))
		case "delete":
			r.router.DELETE(route.Path, wrapHandler(http.HandlerFunc(route.Handler)))
		}
	}
}

//Static add static context to the router
func (r *Router) Static(path string, root string) {
	r.router.ServeFiles(path+"*filepath", http.Dir(root))
}

// Use set the middleware to use by default
func (r *Router) Use(mw ...wrapper.Middleware) {
	for _, m := range mw {
		r.middleware = append(r.middleware, m)
	}
}

//Vars get vars from a request
func (r *Router) Vars(req *http.Request) map[string]string {
	result := map[string]string{}
	ps, ok := req.Context().Value(params).(httprouter.Params)
	if !ok {
		return result
	}
	for _, item := range ps {
		result[item.Key] = item.Value
	}
	log.Debug(result)
	return result
}

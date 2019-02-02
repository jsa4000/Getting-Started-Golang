package pprof

import (
	"net/http/pprof"

	net "webapp/core/net/http"
)

// Controller for http transport
type Controller struct {
	net.RestController
}

// NewController create new RestController
func NewController() net.Controller {
	return &Controller{}
}

// GetRoutes gracefully shutdown rest controller
func (c *Controller) GetRoutes() []net.Route {
	return []net.Route{
		net.Route{
			Path:    "/debug/pprof/",
			Method:  "GET",
			Handler: pprof.Index,
		},
		net.Route{
			Path:    "/debug/pprof/heap",
			Method:  "GET",
			Handler: pprof.Index,
		},
		net.Route{
			Path:    "/debug/pprof/allocs",
			Method:  "GET",
			Handler: pprof.Index,
		},
		net.Route{
			Path:    "/debug/pprof/goroutine",
			Method:  "GET",
			Handler: pprof.Index,
		},
		net.Route{
			Path:    "/debug/pprof/block",
			Method:  "GET",
			Handler: pprof.Index,
		},
		net.Route{
			Path:    "/debug/pprof/mutex",
			Method:  "GET",
			Handler: pprof.Index,
		},
		net.Route{
			Path:    "/debug/pprof/threadcreate",
			Method:  "GET",
			Handler: pprof.Index,
		},
		net.Route{
			Path:    "/debug/pprof/cmdline",
			Method:  "GET",
			Handler: pprof.Cmdline,
		},
		net.Route{
			Path:    "/debug/pprof/profile",
			Method:  "GET",
			Handler: pprof.Profile,
		},
		net.Route{
			Path:    "/debug/pprof/symbol",
			Method:  "GET",
			Handler: pprof.Symbol,
		},
		net.Route{
			Path:    "/debug/pprof/trace",
			Method:  "GET",
			Handler: pprof.Trace,
		},
	}
}

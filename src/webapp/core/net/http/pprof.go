package http

import "net/http/pprof"

// PprofPreffix
const PprofPreffix = "/debug/pprof/"

// PprofController for http transport
type PprofController struct {
	RestController
}

// NewPprofController create new RestController
func NewPprofController() Controller {
	return &RestController{}
}

// GetRoutes gracefully shutdown rest controller
func (c *RestController) GetRoutes() []Route {
	return []Route{
		Route{
			Path:    "/debug/pprof/",
			Method:  "GET",
			Handler: pprof.Index,
		},
		Route{
			Path:    "/debug/pprof/heap",
			Method:  "GET",
			Handler: pprof.Index,
		},
		Route{
			Path:    "/debug/pprof/allocs",
			Method:  "GET",
			Handler: pprof.Index,
		},
		Route{
			Path:    "/debug/pprof/goroutine",
			Method:  "GET",
			Handler: pprof.Index,
		},
		Route{
			Path:    "/debug/pprof/block",
			Method:  "GET",
			Handler: pprof.Index,
		},
		Route{
			Path:    "/debug/pprof/mutex",
			Method:  "GET",
			Handler: pprof.Index,
		},
		Route{
			Path:    "/debug/pprof/threadcreate",
			Method:  "GET",
			Handler: pprof.Index,
		},
		Route{
			Path:    "/debug/pprof/cmdline",
			Method:  "GET",
			Handler: pprof.Cmdline,
		},
		Route{
			Path:    "/debug/pprof/profile",
			Method:  "GET",
			Handler: pprof.Profile,
		},
		Route{
			Path:    "/debug/pprof/symbol",
			Method:  "GET",
			Handler: pprof.Symbol,
		},
		Route{
			Path:    "/debug/pprof/trace",
			Method:  "GET",
			Handler: pprof.Trace,
		},
	}
}

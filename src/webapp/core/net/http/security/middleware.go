package security

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
)

// Middleware  middleware struct
type Middleware struct {
	net.MiddlewareBase
	handlers    []FilterHandler
	exitOnMatch bool
}

// NewMiddleware creation for Auth
func NewMiddleware(handlers []FilterHandler, priority int, exitOnMatch bool) net.Middleware {
	return &Middleware{
		MiddlewareBase: net.MiddlewareBase{
			Hdlr: nil,
			Prio: priority,
		},
		handlers:    handlers,
		exitOnMatch: exitOnMatch,
	}
}

// Handler returns the HandlerMid
func (a *Middleware) Handler() net.HandlerMid {
	return a.handler
}

// FilterHandler decorator (closure)
func (a *Middleware) handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range a.handlers {
			if t, ok := handler.Matches(r.RequestURI); ok {
				if err := handler.Handle(w, r, t); err != nil {
					if a.exitOnMatch {
						a.Error(w, err)
						return
					}
					log.Error(err)
				}
				if a.exitOnMatch {
					break
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

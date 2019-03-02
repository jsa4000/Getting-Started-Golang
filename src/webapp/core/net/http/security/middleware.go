package security

import (
	"net/http"
	net "webapp/core/net/http"
)

// AuthHandler type redefinition
type AuthHandler = FilterHandler

// FilterHandler interface to manage the authorization method
type FilterHandler interface {
	Matcher
	Handle(w http.ResponseWriter, r *http.Request) error
}

// TODO: Order middleware by Priorities

// Middleware  middleware struct
type Middleware struct {
	net.MiddlewareBase
	handlers []FilterHandler
}

// NewMiddleware creation for Auth
func NewMiddleware(handlers []FilterHandler, priority int) net.Middleware {
	return &Middleware{
		MiddlewareBase: net.MiddlewareBase{
			Hdlr: nil,
			Prio: priority,
		},
		handlers: handlers,
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
			if _, ok := handler.Matches(r.RequestURI); ok {
				if err := handler.Handle(w, r); err != nil {
					a.WriteError(w, err)
					return
				}
				break
			}
		}
		next.ServeHTTP(w, r)
	})
}

package security

import (
	"net/http"
	net "webapp/core/net/http"
)

// TODO: Order middleware by Priorities

// Middleware  middleware struct
type Middleware struct {
	net.MiddlewareBase
	handlers []AuthHandler
}

// NewMiddleware creation for Auth
func NewMiddleware(handlers []AuthHandler, priority int) net.Middleware {
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

// AuthHandler decorator (closure)
func (a *Middleware) handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, handler := range a.handlers {
			if net.Matches(net.RemoveParams(r.RequestURI), handler.Targets()) {
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

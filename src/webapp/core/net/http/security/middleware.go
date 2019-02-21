package security

import (
	"net/http"
	net "webapp/core/net/http"
)

// Middleware returns LogginMiddleware struct
type Middleware struct {
	net.MiddlewareBase
	handlers []AuthHandler
}

// NewMiddleware creation for Athentication Middleware
func NewMiddleware(handlers []AuthHandler) net.Middleware {
	return &Middleware{
		MiddlewareBase: net.MiddlewareBase{
			Hdlr: nil,
			Prio: net.PriorityAuthentication,
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
			if net.Contains(r.RequestURI, handler.Targets()) {
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

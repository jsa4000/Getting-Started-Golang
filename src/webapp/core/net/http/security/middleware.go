package security

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
)

// MiddlewareType type of middleware to process
type MiddlewareType int32

const (
	// ExitOnMatch exists middleware on targets match
	ExitOnMatch MiddlewareType = iota
	// ContinueOnMatch continue with other filters even on match
	ContinueOnMatch
)

// AuthHandler type redefinition
type AuthHandler = FilterHandler

// FilterHandler interface to manage the authorization method
type FilterHandler interface {
	Matcher
	Handle(w http.ResponseWriter, r *http.Request, target *Target) error
}

// Middleware  middleware struct
type Middleware struct {
	net.MiddlewareBase
	handlers []FilterHandler
	mtype    MiddlewareType
}

// NewMiddleware creation for Auth
func NewMiddleware(handlers []FilterHandler, priority int, mtype MiddlewareType) net.Middleware {
	return &Middleware{
		MiddlewareBase: net.MiddlewareBase{
			Hdlr: nil,
			Prio: priority,
		},
		handlers: handlers,
		mtype:    mtype,
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
					if a.mtype == ExitOnMatch {
						a.Error(w, err)
						return
					}
					log.Error(err)
				}
				if a.mtype == ExitOnMatch {
					break
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}

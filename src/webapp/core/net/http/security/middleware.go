package security

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/core/errors"
	log "webapp/core/logging"
	net "webapp/core/net/http"
)

var (
	filters = []string{"/swagger/", "/oauth/"}
)

// AuthHandlerMiddleware returns LogginMiddleware struct
type AuthHandlerMiddleware struct {
	net.MiddlewareBase
	//config Config
}

// NewAuthHandlerMiddleware creation
func NewAuthHandlerMiddleware() net.Middleware {
	return &AuthHandlerMiddleware{
		net.MiddlewareBase{
			Hdlr: nil,
			Prio: net.PrioritySecurity,
		},
	}
}

// Handler returns the HandlerMid
func (a *AuthHandlerMiddleware) Handler() net.HandlerMid {
	return a.AuthHandler
}

// AuthHandler decorator (closure)
func (a *AuthHandlerMiddleware) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Verify security for request uri=%s", r.RequestURI)

		if !net.Contains(r.RequestURI, filters) {
			writeError(w, net.ErrForbidden.
				From(fmt.Errorf("Unauthorized for resource '%s'", r.RequestURI)))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// WriteError response
func writeError(w http.ResponseWriter, err error) {
	herr, ok := err.(*errors.Error)
	if !ok {
		herr = net.ErrInternalServer.From(err)
	}
	w.WriteHeader(herr.Code)
	json.NewEncoder(w).Encode(herr)
	log.Error(herr)
}

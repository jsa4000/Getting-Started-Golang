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
	filters = []string{"/swagger/"}
)

// AuthHandlerMiddleware returns LogginMiddleware struct
type AuthHandlerMiddleware struct {
	net.MiddlewareBase
}

// NewAuthHandlerMiddleware creation
func NewAuthHandlerMiddleware() net.Middleware {
	return &AuthHandlerMiddleware{
		net.MiddlewareBase{
			Hdlr: AuthHandler,
			Prio: net.PrioritySecurity,
		},
	}
}

// AuthHandler decorator (closure)
func AuthHandler(next http.Handler) http.Handler {
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

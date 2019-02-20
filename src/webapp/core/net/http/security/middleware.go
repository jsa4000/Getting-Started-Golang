package security

import (
	"net/http"
	net "webapp/core/net/http"
)

const (
	bearerPreffix = "Bearer "
	authHeader    = "Authorization"
)

var (
	noAuth    = []string{"/swagger/"}
	basicAuth = []string{"/oauth/"}
)

// AuthHandlerMiddleware returns LogginMiddleware struct
type AuthHandlerMiddleware struct {
	net.MiddlewareBase
	config *Config
}

// NewAuthHandlerMiddleware creation
func NewAuthHandlerMiddleware(c *Config) net.Middleware {
	return &AuthHandlerMiddleware{
		net.MiddlewareBase{
			Hdlr: nil,
			Prio: net.PrioritySecurity,
		}, c,
	}
}

// Handler returns the HandlerMid
func (a *AuthHandlerMiddleware) Handler() net.HandlerMid {
	return a.AuthHandler
}

// AuthHandler decorator (closure)
func (a *AuthHandlerMiddleware) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if net.Contains(r.RequestURI, basicAuth) {
			if err := BasicAuthHandler(w, r, a.config); err != nil {
				a.WriteError(w, err)
				return
			}
		} else if !net.Contains(r.RequestURI, noAuth) {
			if err := JwtHandler(w, r, a.config); err != nil {
				a.WriteError(w, err)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

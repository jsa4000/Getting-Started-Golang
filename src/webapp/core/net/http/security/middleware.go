package security

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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
	service Service
	config  *Config
}

// NewAuthHandlerMiddleware creation
func NewAuthHandlerMiddleware(c *Config, service Service) net.Middleware {
	return &AuthHandlerMiddleware{
		net.MiddlewareBase{
			Hdlr: nil,
			Prio: net.PrioritySecurity,
		}, service, c,
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
			if err := a.basicAuthHandler(w, r); err != nil {
				a.WriteError(w, err)
				return
			}
		} else if !net.Contains(r.RequestURI, noAuth) {
			if err := a.jwtHandler(w, r); err != nil {
				a.WriteError(w, err)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func (a *AuthHandlerMiddleware) jwtHandler(w http.ResponseWriter, r *http.Request) error {
	basicAuth, ok := r.Header[authHeader]
	if !ok {
		return net.ErrUnauthorized.From(errors.New("Authorization has not been found"))
	}
	token := strings.TrimPrefix(basicAuth[0], bearerPreffix)
	fmt.Println(token)
	resp, err := a.service.CheckToken(r.Context(), &CheckTokenRequest{Token: token})
	if err != nil || !resp.Valid {
		return net.ErrUnauthorized.From(errors.New("Authorization Beared invalid"))
	}
	return nil
}

func (a *AuthHandlerMiddleware) basicAuthHandler(w http.ResponseWriter, r *http.Request) error {
	user, password, hasAuth := r.BasicAuth()
	if !hasAuth {
		return net.ErrUnauthorized.From(errors.New("Authorization has not been found"))
	}
	if user != a.config.ClientID && password != a.config.ClientSecret {
		return net.ErrUnauthorized.From(errors.New("Credentials are not valid for client #{user}"))
	}
	//r.WithContext(context.WithValue(r.Context(), "basicAuth", user))
	return nil
}

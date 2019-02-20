package security

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	net "webapp/core/net/http"
)

// JwtHandler handler to manage token authenticaiton method
func JwtHandler(w http.ResponseWriter, r *http.Request, config *Config) error {
	basicAuth, ok := r.Header[authHeader]
	if !ok {
		return net.ErrUnauthorized.From(errors.New("Authorization has not been found"))
	}
	token := strings.TrimPrefix(basicAuth[0], bearerPreffix)
	fmt.Println(token)
	resp, err := config.Service.Check(r.Context(), &CheckTokenRequest{Token: token})
	if err != nil || !resp.Valid {
		return net.ErrUnauthorized.From(errors.New("Authorization Beared invalid"))
	}
	return nil
}

// BasicAuthHandler handler to manage basic authenticaiton method
func BasicAuthHandler(w http.ResponseWriter, r *http.Request, config *Config) error {
	user, password, hasAuth := r.BasicAuth()
	if !hasAuth {
		return net.ErrUnauthorized.From(errors.New("Authorization has not been found"))
	}
	if user != config.ClientID && password != config.ClientSecret {
		return net.ErrUnauthorized.From(errors.New("Credentials are not valid for client #{user}"))
	}
	//r.WithContext(context.WithValue(r.Context(), "basicAuth", user))
	return nil
}

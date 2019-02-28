package basic

import (
	"errors"
	"fmt"
	"net/http"
	cerr "webapp/core/errors"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

const (
	// AuthKey Key to get data from context in basicAuth
	AuthKey security.ContextKey = "basic-auth-key"
)

// AuthHandler struct to handle basic authentication
type AuthHandler struct {
	*Config
	targets  []string
	local    security.UserInfoService
	provider security.UserInfoService
}

// Handle handler to manage basic authenticaiton method
func (s *AuthHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Basic Auth Request for %s", net.RemoveURLParams(r.RequestURI))
	username, password, hasAuth := r.BasicAuth()
	if !hasAuth {
		return net.ErrUnauthorized.From(errors.New("Authorization is required"))
	}
	var err error
	user, err := s.local.Fetch(r.Context(), username)
	if err != nil {
		herr, ok := err.(*cerr.Error)
		if !ok {
			herr = net.ErrInternalServer.From(err)
		}
		return herr
	}
	if username != user.Name && password != user.Password {
		return net.ErrUnauthorized.From(fmt.Errorf("Credentials are not valid for client %s", user))
	}
	security.SetContextValue(r, AuthKey, new(security.ContextValue))
	security.SetUserName(r, username)
	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *AuthHandler) Targets() []string {
	return s.targets
}

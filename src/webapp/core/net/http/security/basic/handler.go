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
	*security.Targets
	service security.UserInfoService
}

// Handle handler to manage basic authentication method
func (s *AuthHandler) Handle(w http.ResponseWriter, r *http.Request, target *security.Target) error {
	log.Debugf("Handle Basic Auth Request for %s", net.RemoveURLParams(r.RequestURI))
	username, password, hasAuth := r.BasicAuth()
	if !hasAuth {
		return net.ErrUnauthorized.From(errors.New("Authorization is required"))
	}
	var err error
	user, err := s.service.Fetch(r.Context(), username)
	if err != nil {
		herr, ok := err.(*cerr.Error)
		if !ok {
			herr = net.ErrInternalServer.From(err)
		}
		return herr
	}
	if !security.ValidateUser(user, username, password) {
		return net.ErrUnauthorized.From(fmt.Errorf("Invalid credentials for user %s", user))
	}
	if !target.Any(user.Roles) {
		return net.ErrForbidden.From(fmt.Errorf("User %s has not enough privileges", user))
	}
	security.SetContextValue(r, AuthKey, new(security.ContextValue))
	security.SetUserName(r, username)
	security.SetUserID(r, user.ID)
	security.SetUserRoles(r, user.Roles)
	return nil
}

package basic

import (
	"errors"
	"fmt"
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

const (
	// ContextValue Key to get data from context in basicAuth
	ContextValue string = "basic-auth-key"
)

// Handler struct to handle basic authentication
type Handler struct {
	security.BaseHandler
	*Config
	service security.UserInfoService
}

// Handle handler to manage basic authentication method
func (s *Handler) Handle(w http.ResponseWriter, r *http.Request, target security.Target) error {
	log.Debugf("Handle Basic Auth Request for %s", net.RemoveURLParams(r.RequestURI))
	username, password, hasAuth := r.BasicAuth()
	if !hasAuth {
		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return net.ErrUnauthorized.From(errors.New("Authorization is required"))
	}
	var err error
	user, err := s.service.Fetch(r.Context(), username)
	if err != nil {
		return err
	}
	if !security.ValidateUser(user, username, password) {
		return net.ErrUnauthorized.From(fmt.Errorf("Invalid credentials for user %s", user))
	}
	if !target.Any(user.Roles) {
		return net.ErrForbidden.From(fmt.Errorf("User %s has not enough privileges", user))
	}
	security.SetAuthType(r, ContextValue)
	security.SetUserName(r, username)
	security.SetUserID(r, user.ID)
	security.SetUserRoles(r, user.Roles)
	return nil
}

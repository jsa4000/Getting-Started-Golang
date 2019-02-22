package basic

import (
	"errors"
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

const (
	// AuthKey Key to get data from context in basicAuth
	AuthKey security.ContextKey = "basic-auth-key"
)

// Service struct to handle basic authentication
type Service struct {
	*Config
	targets     []string
	userFetcher security.UserFetcher
}

// Handle handler to manage basic authenticaiton method
func (s *Service) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Basic Auth Request for %s", net.RemoveParams(r.RequestURI))
	user, password, hasAuth := r.BasicAuth()
	if !hasAuth {
		return net.ErrUnauthorized.From(errors.New("Authorization has not been found"))
	}
	if user != s.Config.ClientID && password != s.Config.ClientSecret {
		return net.ErrUnauthorized.From(errors.New("Credentials are not valid for client #{user}"))
	}
	security.SetContextValue(r, AuthKey, new(security.ContextValue))
	security.SetUserName(r, user)
	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *Service) Targets() []string {
	return s.targets
}

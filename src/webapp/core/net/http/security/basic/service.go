package basic

import (
	"errors"
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// Service struct to handle basic authentication
type Service struct {
	*Config
	targets     []string
	userFetcher security.UserFetcher
}

// Handle handler to manage basic authenticaiton method
func (s *Service) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Basic Auth Request for %s", r.RequestURI)
	user, password, hasAuth := r.BasicAuth()
	if !hasAuth {
		return net.ErrUnauthorized.From(errors.New("Authorization has not been found"))
	}
	if user != s.Config.ClientID && password != s.Config.ClientSecret {
		return net.ErrUnauthorized.From(errors.New("Credentials are not valid for client #{user}"))
	}
	//r.WithContext(context.WithValue(r.Context(), "basicAuth", user))
	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *Service) Targets() []string {
	return s.targets
}

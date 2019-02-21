package role

import (
	"net/http"
	log "webapp/core/logging"
	"webapp/core/net/http/security"
)

type key string

const (
	// BasicAuthKey Key to get data from context in basicAuth
	BasicAuthKey key = "BasicAuthKey"
)

// Service struct to handle basic authentication
type Service struct {
	*Config
	targets     []string
	userFetcher security.UserFetcher
}

// Handle handler to manage basic authenticaiton method
func (s *Service) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Role Based Authorization Request for %s", r.RequestURI)

	//value, ok := r.Context().Value( BasicAuthKey, user))

	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *Service) Targets() []string {
	return s.targets
}

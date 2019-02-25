package oauth

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

const (
	// AuthKey Key to get data from context in basicAuth
	AuthKey security.ContextKey = "oauth-auth-key"
)

// Service struct to handle basic authentication
type Service struct {
	*Config
	targets []string
}

// Handle handler to manage Oauth authenticaiton method
func (s *Service) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Oauth Auth Request for %s", net.RemoveParams(r.RequestURI))

	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *Service) Targets() []string {
	return s.targets
}

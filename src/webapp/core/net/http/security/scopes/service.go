package scopes

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// Service struct to handle resource authorization by scopes
type Service struct {
	*Config
	targets []string
}

// Handle handler to manage scopes authorization method
func (s *Service) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Role Based Authorization Request for %s", net.RemoveParams(r.RequestURI))
	if value := r.Context().Value(security.UserRolesKey); value != nil {
		log.Debugf("Roles %v", value)
	}
	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *Service) Targets() []string {
	return s.targets
}

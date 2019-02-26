package scopes

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// AuthHandler struct to handle resource authorization by scopes
type AuthHandler struct {
	*Config
	targets []string
}

// Handle handler to manage scopes authorization method
func (s *AuthHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Role Based Authorization Request for %s", net.RemoveURLParams(r.RequestURI))
	if value := r.Context().Value(security.UserRolesKey); value != nil {
		log.Debugf("Roles %v", value)
	}
	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *AuthHandler) Targets() []string {
	return s.targets
}

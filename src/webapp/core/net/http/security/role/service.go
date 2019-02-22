package role

import (
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
	log.Debugf("Handle Role Based Authorization Request for %s", net.RemoveParams(r.RequestURI))

	if value := r.Context().Value(security.UserNameKey); value != nil {
		log.Debugf("UserName %v", value)
	}

	if value := r.Context().Value(security.UserIDKey); value != nil {
		log.Debugf("UserID %v", value)
	}

	if value := r.Context().Value(security.UserRolesKey); value != nil {
		log.Debugf("Roles %v", value)
	}

	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *Service) Targets() []string {
	return s.targets
}

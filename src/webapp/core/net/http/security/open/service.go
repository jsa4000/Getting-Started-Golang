package open

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
)

// Service struct to handle basic authentication
type Service struct {
	*Config
	targets []string
}

// Handle handler to manage basic authenticaiton method
func (s *Service) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Open Request for %s", net.RemoveParams(r.RequestURI))
	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *Service) Targets() []string {
	return s.targets
}

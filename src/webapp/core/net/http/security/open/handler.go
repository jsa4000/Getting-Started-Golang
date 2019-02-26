package open

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
)

// AuthHandler struct to handle basic authentication
type AuthHandler struct {
	*Config
	targets []string
}

// Handle handler to manage basic authenticaiton method
func (s *AuthHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Open Request for %s", net.RemoveURLParams(r.RequestURI))
	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *AuthHandler) Targets() []string {
	return s.targets
}

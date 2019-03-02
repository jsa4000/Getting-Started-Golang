package open

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// AuthHandler struct to handle basic authentication
type AuthHandler struct {
	*Config
	*security.Targets
}

// Handle handler to manage basic authenticaiton method
func (s *AuthHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle Open Request for %s", net.RemoveURLParams(r.RequestURI))
	return nil
}

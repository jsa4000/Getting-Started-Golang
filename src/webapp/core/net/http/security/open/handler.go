package open

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// Handler struct to handle basic authentication
type Handler struct {
	security.BaseHandler
	*Config
}

// Handle handler to manage basic authenticaiton method
func (s *Handler) Handle(w http.ResponseWriter, r *http.Request, target security.Target) error {
	log.Debugf("Handle Open Request for %s", net.RemoveURLParams(r.RequestURI))
	return nil
}

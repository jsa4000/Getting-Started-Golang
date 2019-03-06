package roles

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
	auth, err := security.ContextValue(r)
	if err != nil {
		return nil
	}
	log.Debugf("Handle Roles Request for %s and auth type %s", net.RemoveURLParams(r.RequestURI), auth)
	return nil
}

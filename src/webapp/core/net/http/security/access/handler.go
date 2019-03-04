package access

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// AuthHandler struct to handle access control methods
type AuthHandler struct {
	security.BaseHandler
	*Config
}

// Handle handler to manage access control methods
func (s *AuthHandler) Handle(w http.ResponseWriter, r *http.Request, target security.Target) error {
	log.Debugf("Handle Access Request for %s", net.RemoveURLParams(r.RequestURI))
	if access, ok := target.(*Target); ok && access.Allow {
		enableCors(w, access.Origin)
	}
	return nil
}

func defaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func enableCors(w http.ResponseWriter, origin string) {
	w.Header().Set("Access-Control-Allow-Origin", origin)
}
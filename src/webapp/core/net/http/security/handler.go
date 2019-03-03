package security

import (
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
)

// AuthHandler type redefinition
type AuthHandler = FilterHandler

// FilterHandler interface to manage the authorization method
type FilterHandler interface {
	Matches(url string) (Target, bool)
	Handle(w http.ResponseWriter, r *http.Request, target Target) error
}

// BaseHandler struct to handle access control methods
type BaseHandler struct {
	Targets
}

// Matches any target with the given URL
func (b *BaseHandler) Matches(url string) (Target, bool) {
	for _, t := range b.Targets {
		if t.Matches(url) {
			return t, true
		}
	}
	return nil, false
}

// Handle handler to manage access control methods
func (b *BaseHandler) Handle(w http.ResponseWriter, r *http.Request, target Target) error {
	log.Debugf("Handle Request for %s", net.RemoveURLParams(r.RequestURI))
	return nil
}

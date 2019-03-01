package security

import (
	"net/http"
)

// AuthHandler type redefinition
type AuthHandler = FilterHandler

// FilterHandler interface to manage the authorization method
type FilterHandler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
	Targets() *Targets
}

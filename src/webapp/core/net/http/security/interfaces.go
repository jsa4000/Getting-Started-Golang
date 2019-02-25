package security

import (
	"net/http"
)

// AuthHandler interface to manage the authorization method
type AuthHandler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
	Targets() []string
}

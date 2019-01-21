package transport

import "net/http"

// HTTPHandler for handle the requests
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// HTTPRoute route
type HTTPRoute struct {
	Path    string
	Method  string
	Handler HTTPHandler
	secured bool
	roles   []string
}

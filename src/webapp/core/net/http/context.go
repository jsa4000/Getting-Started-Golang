package http

import (
	"net/http"
)

// ContextKey type for context keys
type ContextKey string

const (
	// RouteInfoKey Key used to extract rout information for filters
	RouteInfoKey ContextKey = "route"
)

// RouteInfo set user name into Context
func RouteInfo(r *http.Request) (*Route, bool) {
	result, ok := r.Context().Value(RouteInfoKey).(Route)
	if !ok {
		return nil, false
	}
	return &result, true
}

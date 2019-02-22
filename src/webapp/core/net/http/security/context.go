package security

import (
	"context"
	"net/http"
)

// ContextKey type for context keys
type ContextKey string

// ContextValue type for context values
type ContextValue struct{}

const (
	// UserNameKey Key to get the user (name or email) data from context during the authorization
	UserNameKey ContextKey = "user-name-key"
	// UserIDKey Key to get the internal userID from context during the authorization
	UserIDKey ContextKey = "user-id-key"
	// UserRolesKey Key to get user roles from context during the authorization
	UserRolesKey ContextKey = "user-roles-key"
)

// SetContextValue set user name into Context
func SetContextValue(r *http.Request, key ContextKey, value interface{}) {
	*r = *r.WithContext(context.WithValue(r.Context(), key, value))
}

// SetUserName set user name into Context
func SetUserName(r *http.Request, username string) {
	*r = *r.WithContext(context.WithValue(r.Context(), UserNameKey, username))
}

// SetUserID set user id into Context
func SetUserID(r *http.Request, id string) {
	*r = *r.WithContext(context.WithValue(r.Context(), UserIDKey, id))
}

// SetRoles set user id into Context
func SetRoles(r *http.Request, roles []string) {
	*r = *r.WithContext(context.WithValue(r.Context(), UserRolesKey, roles))
}

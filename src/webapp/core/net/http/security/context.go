package security

import (
	"context"
	"net/http"
)

// ContextKey type for context keys
type ContextKey string

const (
	// AuthKey Key to get the user (name or email) data from context during the authorization
	AuthKey ContextKey = "auth-key"
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

// AuthType set user name into Context
func AuthType(r *http.Request) (string, bool) {
	result, ok := r.Context().Value(AuthKey).(string)
	if !ok {
		return "", false
	}
	return result, true
}

// SetAuthType set user name into Context
func SetAuthType(r *http.Request, value string) {
	*r = *r.WithContext(context.WithValue(r.Context(), AuthKey, value))
}

// GetUserName set user id into Context
func GetUserName(r *http.Request) (string, bool) {
	result, ok := r.Context().Value(UserNameKey).(string)
	if !ok {
		return "", false
	}
	return result, true
}

// SetUserName set user name into Context
func SetUserName(r *http.Request, username string) {
	*r = *r.WithContext(context.WithValue(r.Context(), UserNameKey, username))
}

// SetUserID set user id into Context
func SetUserID(r *http.Request, id string) {
	*r = *r.WithContext(context.WithValue(r.Context(), UserIDKey, id))
}

// UserRoles set user id into Context
func UserRoles(r *http.Request) ([]string, bool) {
	result, ok := r.Context().Value(UserRolesKey).([]string)
	if !ok {
		return nil, false
	}
	return result, true
}

// SetUserRoles set user id into Context
func SetUserRoles(r *http.Request, roles []string) {
	*r = *r.WithContext(context.WithValue(r.Context(), UserRolesKey, roles))
}

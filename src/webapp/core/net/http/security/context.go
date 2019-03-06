package security

import (
	"context"
	"errors"
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

// ContextValue set user name into Context
func ContextValue(r *http.Request) (string, error) {
	result := r.Context().Value(AuthKey)
	if result == nil {
		return "", errors.New("Error Getting the value from context")
	}
	return result.(string), nil
}

// SetContextValue set user name into Context
func SetContextValue(r *http.Request, key ContextKey, value interface{}) {
	*r = *r.WithContext(context.WithValue(r.Context(), key, value))
}

// SetAuthKey set user name into Context
func SetAuthKey(r *http.Request, value string) {
	*r = *r.WithContext(context.WithValue(r.Context(), AuthKey, value))
}

// SetUserName set user name into Context
func SetUserName(r *http.Request, username string) {
	*r = *r.WithContext(context.WithValue(r.Context(), UserNameKey, username))
}

// SetUserID set user id into Context
func SetUserID(r *http.Request, id string) {
	*r = *r.WithContext(context.WithValue(r.Context(), UserIDKey, id))
}

// SettUsersRoles set user id into Context
func SetUserRoles(r *http.Request, roles []string) {
	*r = *r.WithContext(context.WithValue(r.Context(), UserRolesKey, roles))
}

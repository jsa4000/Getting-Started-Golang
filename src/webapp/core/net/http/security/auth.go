package security

import (
	"context"
	"fmt"
	net "webapp/core/net/http"
)

// AuthManager struct to handle basic authentication
type AuthManager struct {
	services []UserInfoService
}

// AuthManagerBuilder struct to handle basic authentication
type AuthManagerBuilder struct {
	*AuthManager
	authUsersBuilder *NestedAuthUsersBuilder
}

// NewAuthManagerBuilder Create a new AuthUsersBuilder
func NewAuthManagerBuilder() *AuthManagerBuilder {
	return &AuthManagerBuilder{
		AuthManager: &AuthManager{
			services: make([]UserInfoService, 0),
		},
	}
}

// NestedAuthUsersBuilder build struct
type NestedAuthUsersBuilder struct {
	*AuthUsersBuilder
	parent *AuthManagerBuilder
}

// NewAuthUsersBuilder Create a new AuthUsersBuilder
func newNestedAuthUsersBuilder(parent *AuthManagerBuilder) *NestedAuthUsersBuilder {
	return &NestedAuthUsersBuilder{
		NewAuthUsersBuilder(),
		parent,
	}
}

// Fetch Add In memory user service
// It uses the in-memory used first to search for
func (am *AuthManager) Fetch(ctx context.Context, username string) (*UserInfo, error) {
	for _, s := range am.services {
		if user, err := s.Fetch(ctx, username); err == nil {
			return user, nil
		}
	}
	return nil, net.ErrNotFound.From(fmt.Errorf("User %s has not been found", username))
}

// WithUserService Add user service
func (b *AuthManagerBuilder) WithUserService(s UserInfoService) *AuthManagerBuilder {
	b.services = append(b.services, s)
	return b
}

// WithInMemoryUsers set the interface to use for fetching user info
func (b *AuthManagerBuilder) WithInMemoryUsers() *NestedAuthUsersBuilder {
	if b.authUsersBuilder == nil {
		b.authUsersBuilder = newNestedAuthUsersBuilder(b)
	}
	return b.authUsersBuilder
}

// Build Add user service
func (b *AuthManagerBuilder) Build() *AuthManager {
	b.services = append([]UserInfoService{b.authUsersBuilder.Build()}, b.services...)
	return b.AuthManager
}

// WithUser set the interface to use for fetching user info
func (c *NestedAuthUsersBuilder) WithUser(name string) *NestedAuthUsersBuilder {
	c.AuthUsersBuilder.WithUser(name)
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *NestedAuthUsersBuilder) WithPassword(password string) *NestedAuthUsersBuilder {
	c.AuthUsersBuilder.WithPassword(password)
	return c
}

// WithRoles set the interface to use for fetching user info
func (c *NestedAuthUsersBuilder) WithRoles(roles []string) *NestedAuthUsersBuilder {
	c.AuthUsersBuilder.WithRoles(roles)
	return c
}

// And set the interface to use for fetching user info
func (c *NestedAuthUsersBuilder) And() *AuthManagerBuilder {
	return c.parent
}

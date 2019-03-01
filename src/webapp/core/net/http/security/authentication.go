package security

import (
	"context"
	"fmt"
	net "webapp/core/net/http"
)

// AuthenticationManager struct to handle basic authentication
type AuthenticationManager struct {
	services []UserInfoService
}

// AuthenticationManagerBuilder struct to handle basic authentication
type AuthenticationManagerBuilder struct {
	*AuthenticationManager
	authUsersBuilder *NestedAuthorizedUsersBuilder
}

// NewAuthenticationManagerBuilder Create a new AuthorizedUsersBuilder
func NewAuthenticationManagerBuilder() *AuthenticationManagerBuilder {
	return &AuthenticationManagerBuilder{
		AuthenticationManager: &AuthenticationManager{
			services: make([]UserInfoService, 0),
		},
	}
}

// NestedAuthorizedUsersBuilder build struct
type NestedAuthorizedUsersBuilder struct {
	*AuthorizedUsersBuilder
	parent *AuthenticationManagerBuilder
}

// NewAuthorizedUsersBuilder Create a new AuthorizedUsersBuilder
func newNestedAuthorizedUsersBuilder(parent *AuthenticationManagerBuilder) *NestedAuthorizedUsersBuilder {
	return &NestedAuthorizedUsersBuilder{
		NewAuthorizedUsersBuilder(),
		parent,
	}
}

// Fetch Add In memory user service
// It uses the in-memory used first to search for
func (am *AuthenticationManager) Fetch(ctx context.Context, username string) (*UserInfo, error) {
	for _, s := range am.services {
		if user, err := s.Fetch(ctx, username); err == nil {
			return user, nil
		}
	}
	return nil, net.ErrNotFound.From(fmt.Errorf("User %s has not been found", username))
}

// WithUserService Add user service
func (b *AuthenticationManagerBuilder) WithUserService(s UserInfoService) *AuthenticationManagerBuilder {
	b.services = append(b.services, s)
	return b
}

// WithInMemoryUsers set the interface to use for fetching user info
func (b *AuthenticationManagerBuilder) WithInMemoryUsers() *NestedAuthorizedUsersBuilder {
	if b.authUsersBuilder == nil {
		b.authUsersBuilder = newNestedAuthorizedUsersBuilder(b)
	}
	return b.authUsersBuilder
}

// Build Add user service
func (b *AuthenticationManagerBuilder) Build() *AuthenticationManager {
	b.services = append([]UserInfoService{b.authUsersBuilder.Build()}, b.services...)
	return b.AuthenticationManager
}

// WithUser set the interface to use for fetching user info
func (c *NestedAuthorizedUsersBuilder) WithUser(name string) *NestedAuthorizedUsersBuilder {
	c.AuthorizedUsersBuilder.WithUser(name)
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *NestedAuthorizedUsersBuilder) WithPassword(password string) *NestedAuthorizedUsersBuilder {
	c.AuthorizedUsersBuilder.WithPassword(password)
	return c
}

// WithRoles set the interface to use for fetching user info
func (c *NestedAuthorizedUsersBuilder) WithRoles(roles []string) *NestedAuthorizedUsersBuilder {
	c.AuthorizedUsersBuilder.WithRoles(roles)
	return c
}

// And set the interface to use for fetching user info
func (c *NestedAuthorizedUsersBuilder) And() *AuthenticationManagerBuilder {
	return c.parent
}

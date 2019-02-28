package security

import (
	"context"
	"fmt"
	net "webapp/core/net/http"
)

// AuthorizedUsers to implement the UserinfoProvider
type AuthorizedUsers struct {
	Users map[string]*UserInfo
}

// Fetch LocalUserProvider UserInfoService interface
func (c *AuthorizedUsers) Fetch(ctx context.Context, username string) (*UserInfo, error) {
	user, ok := c.Users[username]
	if !ok {
		return nil, net.ErrNotFound.From(fmt.Errorf("User %s has not been found", username))
	}
	return user, nil
}

// AuthorizedUsersBuilder build struct
type AuthorizedUsersBuilder struct {
	*AuthorizedUsers
	current *UserInfo
}

// NewAuthorizedUsersBuilder Create a new DefaultUsersBuilder
func NewAuthorizedUsersBuilder() *AuthorizedUsersBuilder {
	return &AuthorizedUsersBuilder{
		AuthorizedUsers: &AuthorizedUsers{
			Users: make(map[string]*UserInfo),
		},
	}
}

// WithUser set the interface to use for fetching user info
func (c *AuthorizedUsersBuilder) WithUser(name string) *AuthorizedUsersBuilder {
	if c.current != nil {
		c.Users[c.current.Name] = c.current
	}
	c.current = &UserInfo{
		Name: name,
	}
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *AuthorizedUsersBuilder) WithPassword(password string) *AuthorizedUsersBuilder {
	c.current.Password = password
	return c
}

// WithRoles set the interface to use for fetching user info
func (c *AuthorizedUsersBuilder) WithRoles(roles []string) *AuthorizedUsersBuilder {
	c.current.Roles = roles
	return c
}

// Build default users struct
func (c *AuthorizedUsersBuilder) Build() *AuthorizedUsers {
	if c.current != nil {
		c.Users[c.current.Name] = c.current
	}
	return c.AuthorizedUsers
}

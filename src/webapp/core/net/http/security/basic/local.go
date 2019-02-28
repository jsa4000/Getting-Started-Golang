package basic

import (
	"context"
	"fmt"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// DefaultUsers to implement the UserinfoProvider
type DefaultUsers struct {
	Users map[string]*security.UserInfo
}

// Fetch LocalUserProvider UserInfoProvider interface
func (c *DefaultUsers) Fetch(ctx context.Context, username string) (*security.UserInfo, error) {
	user, ok := c.Users[username]
	if !ok {
		return nil, net.ErrNotFound.From(fmt.Errorf("User %s has not been found", username))
	}
	return user, nil
}

// DefaultUsersBuilder build struct
type DefaultUsersBuilder struct {
	*DefaultUsers
	current *security.UserInfo
}

// NewDefaultUsersBuilder Create a new DefaultUsersBuilder
func NewDefaultUsersBuilder() *DefaultUsersBuilder {
	return &DefaultUsersBuilder{
		DefaultUsers: &DefaultUsers{
			Users: make(map[string]*security.UserInfo),
		},
	}
}

// WithUser set the interface to use for fetching user info
func (c *DefaultUsersBuilder) WithUser(name string) *DefaultUsersBuilder {
	if c.current != nil {
		c.Users[c.current.Name] = c.current
	}
	c.current = &security.UserInfo{
		Name: name,
	}
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *DefaultUsersBuilder) WithPassword(password string) *DefaultUsersBuilder {
	c.current.Password = password
	return c
}

// WithRoles set the interface to use for fetching user info
func (c *DefaultUsersBuilder) WithRoles(roles []string) *DefaultUsersBuilder {
	c.current.Roles = roles
	return c
}

// Build default users struct
func (c *DefaultUsersBuilder) Build() *DefaultUsers {
	if c.current != nil {
		c.Users[c.current.Name] = c.current
	}
	return c.DefaultUsers
}

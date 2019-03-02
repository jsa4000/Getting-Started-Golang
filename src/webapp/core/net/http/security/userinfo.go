package security

import (
	"context"
	"fmt"
	pcrypt "webapp/core/crypto/password"
	net "webapp/core/net/http"
)

// UserInfo structure to retrieve (fetch) the user information
type UserInfo struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Password  string                 `json:"password"`
	Roles     []string               `json:"roles"`
	Resources map[string]interface{} `json:"resources"`
}

// UserInfoService Interface
type UserInfoService interface {
	Fetch(ctx context.Context, username string) (*UserInfo, error)
}

// AuthUsers to implement the UserinfoProvider
type AuthUsers struct {
	Users map[string]*UserInfo
}

// Fetch LocalUserProvider UserInfoService interface
func (c *AuthUsers) Fetch(ctx context.Context, username string) (*UserInfo, error) {
	user, ok := c.Users[username]
	if !ok {
		return nil, net.ErrNotFound.From(fmt.Errorf("User %s has not been found", username))
	}
	return user, nil
}

// AuthUsersBuilder build struct
type AuthUsersBuilder struct {
	*AuthUsers
	current *UserInfo
}

// NewAuthUsersBuilder Create a new DefaultUsersBuilder
func NewAuthUsersBuilder() *AuthUsersBuilder {
	return &AuthUsersBuilder{
		AuthUsers: &AuthUsers{
			Users: make(map[string]*UserInfo),
		},
	}
}

// WithUser set the interface to use for fetching user info
func (c *AuthUsersBuilder) WithUser(name string) *AuthUsersBuilder {
	if c.current != nil {
		c.Users[c.current.Name] = c.current
	}
	c.current = &UserInfo{
		Name: name,
	}
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *AuthUsersBuilder) WithPassword(password string) *AuthUsersBuilder {
	c.current.Password = pcrypt.New(password)
	return c
}

// WithRoles set the interface to use for fetching user info
func (c *AuthUsersBuilder) WithRoles(roles []string) *AuthUsersBuilder {
	c.current.Roles = roles
	return c
}

// Build default users struct
func (c *AuthUsersBuilder) Build() *AuthUsers {
	if c.current != nil {
		c.Users[c.current.Name] = c.current
	}
	return c.AuthUsers
}

// ValidateUser compares the username and password to be the same as the userinfo
func ValidateUser(user *UserInfo, username, password string) bool {
	return (username == user.Name || username == user.Email) &&
		pcrypt.Compare(user.Password, password)
}

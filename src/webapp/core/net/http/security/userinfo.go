package security

import (
	"context"
	"fmt"
	"strings"
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

// Matches retrieves the roles of the user with the requested scopes
func (user *UserInfo) Matches(roles []string) []string {
	result := make([]string, 0)
	if len(user.Roles) == 0 {
		return result
	}
	for _, s := range roles {
		s = strings.ToLower(s)
		for _, cs := range user.Roles {
			if s == strings.ToLower(cs) {
				result = append(result, s)
				continue
			}
		}
	}
	return result
}

// UserInfoService Interface
type UserInfoService interface {
	Fetch(ctx context.Context, username string) (*UserInfo, error)
}

// ValidateUser compares the username and password to be the same as the userinfo
func ValidateUser(user *UserInfo, username, password string) bool {
	return (username == user.Name || username == user.Email) &&
		pcrypt.Compare(user.Password, password)
}

// Users to implement the UserinfoProvider
type Users map[string]*UserInfo

// Fetch LocalUserProvider UserInfoService interface
func (c Users) Fetch(ctx context.Context, username string) (*UserInfo, error) {
	user, ok := c[username]
	if !ok {
		return nil, net.ErrNotFound.From(fmt.Errorf("User %s has not been found", username))
	}
	return user, nil
}

// UsersBuilder build struct
type UsersBuilder struct {
	Users
	current *UserInfo
}

// NewUsersBuilder Create a new DefaultUsersBuilder
func NewUsersBuilder() *UsersBuilder {
	return &UsersBuilder{
		Users: make(map[string]*UserInfo),
	}
}

// WithUser set the interface to use for fetching user info
func (c *UsersBuilder) WithUser(name string) *UsersBuilder {
	if c.current != nil {
		c.Users[c.current.Name] = c.current
	}
	c.current = &UserInfo{
		Name: name,
	}
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *UsersBuilder) WithPassword(password string) *UsersBuilder {
	c.current.Password = pcrypt.New(password)
	return c
}

// WithRole set the interface to use for fetching user info
func (c *UsersBuilder) WithRole(roles ...string) *UsersBuilder {
	c.current.Roles = append(c.current.Roles, roles...)
	return c
}

// Build default users struct
func (c *UsersBuilder) Build() Users {
	if c.current != nil {
		c.Users[c.current.Name] = c.current
	}
	return c.Users
}

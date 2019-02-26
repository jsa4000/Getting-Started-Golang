package security

import (
	"context"
	"fmt"
	net "webapp/core/net/http"

	uuid "github.com/satori/go.uuid"
)

// UserProvider to implement the UserinfoProvider
type UserProvider struct {
	Users map[string]UserInfo
}

// Fetch implements UserInfoProvider interface
func (c *UserProvider) Fetch(ctx context.Context, username string) (*UserInfo, error) {
	user, ok := c.Users[username]
	if !ok {
		return nil, net.ErrNotFound.From(fmt.Errorf("User %s has not been found", username))
	}
	return &user, nil
}

// UserProviderBuilder to implement the UserinfoProvider
type UserProviderBuilder struct {
	*UserProvider
}

// NewUserProviderBuilder Create a Mock repository
func NewUserProviderBuilder() *UserProviderBuilder {
	return &UserProviderBuilder{
		&UserProvider{
			Users: make(map[string]UserInfo),
		},
	}
}

// WithUser Create a Mock repository
func (u *UserProviderBuilder) WithUser(username string, password string) *UserProviderBuilder {
	u.UserProvider.Users[username] = UserInfo{
		ID:       uuid.NewV4().String(),
		Name:     username,
		Password: password,
	}
	return u
}

// Build Create UserProvider
func (u *UserProviderBuilder) Build() *UserProvider {
	return u.UserProvider
}

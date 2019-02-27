package basic

import (
	"context"
	"fmt"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// LocalUsers to implement the UserinfoProvider
type LocalUsers struct {
	Users map[string]*security.UserInfo
}

// NewlocalUsers Return new instance of local users
func NewlocalUsers() *LocalUsers {
	return &LocalUsers{
		Users: make(map[string]*security.UserInfo),
	}
}

// Fetch LocalUserProvider UserInfoProvider interface
func (c *LocalUsers) Fetch(ctx context.Context, username string) (*security.UserInfo, error) {
	user, ok := c.Users[username]
	if !ok {
		return nil, net.ErrNotFound.From(fmt.Errorf("User %s has not been found", username))
	}
	return user, nil
}

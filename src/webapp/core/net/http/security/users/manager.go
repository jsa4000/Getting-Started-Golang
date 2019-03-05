package users

import (
	"context"
	"fmt"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// Manager struct to handle basic authentication
type Manager struct {
	services []security.UserInfoService
}

// Fetch Add In memory user service
// It uses the in-memory used first to search for
func (am *Manager) Fetch(ctx context.Context, username string) (*security.UserInfo, error) {
	for _, s := range am.services {
		if user, err := s.Fetch(ctx, username); err == nil {
			return user, nil
		}
	}
	return nil, net.ErrNotFound.From(fmt.Errorf("User %s has not been found", username))
}

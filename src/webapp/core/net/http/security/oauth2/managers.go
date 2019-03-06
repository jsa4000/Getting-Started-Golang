package oauth2

import (
	"context"
	"fmt"
	net "webapp/core/net/http"
	"webapp/core/net/http/security/token"
)

// Manager struct to handle basic authentication
type Manager struct {
	clientmanager *ClientManager
	tokenService  token.Service
	controller    net.Controller
}

// Controllers returns the controller for the security implementation
func (m *Manager) Controllers() []net.Controller {
	return []net.Controller{m.controller}
}

// ClientManager struct to handle basic authentication
type ClientManager struct {
	clientServices []ClientService
}

// Fetch Add In memory user service
// It uses the in-memory used first to search for
func (m *ClientManager) Fetch(ctx context.Context, username string) (*Client, error) {
	for _, s := range m.clientServices {
		if user, err := s.Fetch(ctx, username); err == nil {
			return user, nil
		}
	}
	return nil, net.ErrNotFound.From(fmt.Errorf("Client %s has not been found", username))
}

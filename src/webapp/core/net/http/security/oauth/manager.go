package oauth

import (
	"context"
	"fmt"
	net "webapp/core/net/http"
	"webapp/core/net/http/security/token"
)

// Manager struct to handle basic authentication
type Manager struct {
	services     []ClientService
	tokenService token.Service
	controller   net.Controller
}

// Controllers returns the controller for the security implementation
func (m *Manager) Controllers() []net.Controller {
	return []net.Controller{m.controller}
}

// Fetch Add In memory user service
// It uses the in-memory used first to search for
func (m *Manager) Fetch(ctx context.Context, username string) (*Client, error) {
	for _, s := range m.services {
		if user, err := s.Fetch(ctx, username); err == nil {
			return user, nil
		}
	}
	return nil, net.ErrNotFound.From(fmt.Errorf("Client %s has not been found", username))
}

// ManagerBuilder struct to handle basic authentication
type ManagerBuilder struct {
	*Manager
	authClientsBuilder *NestedClientsBuilder
}

// NewManagerBuilder Create a new ClientsBuilder
func NewManagerBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		Manager: &Manager{
			services: make([]ClientService, 0),
		},
	}
}

// NestedClientsBuilder build struct
type NestedClientsBuilder struct {
	*ClientsBuilder
	parent *ManagerBuilder
}

// NewClientsBuilder Create a new ClientsBuilder
func newNestedClientsBuilder(parent *ManagerBuilder) *NestedClientsBuilder {
	return &NestedClientsBuilder{
		NewClientsBuilder(),
		parent,
	}
}

// WithClientService Add user service
func (b *ManagerBuilder) WithClientService(s ClientService) *ManagerBuilder {
	b.services = append(b.services, s)
	return b
}

// WithTokenService set the interface to use for fetch user info
func (b *ManagerBuilder) WithTokenService(ts token.Service) *ManagerBuilder {
	b.tokenService = ts
	return b
}

// WithInMemoryClients set the interface to use for fetching user info
func (b *ManagerBuilder) WithInMemoryClients() *NestedClientsBuilder {
	if b.authClientsBuilder == nil {
		b.authClientsBuilder = newNestedClientsBuilder(b)
	}
	return b.authClientsBuilder
}

// Build Add user service
func (b *ManagerBuilder) Build() *Manager {
	b.controller = NewRestController(b.tokenService)
	b.services = append([]ClientService{b.authClientsBuilder.Build()}, b.services...)
	return b.Manager
}

// WithClient set the interface to use for fetching user info
func (c *NestedClientsBuilder) WithClient(name string) *NestedClientsBuilder {
	c.ClientsBuilder.WithClient(name)
	return c
}

// WithSecret set the interface to use for fetching user info
func (c *NestedClientsBuilder) WithSecret(secret string) *NestedClientsBuilder {
	c.ClientsBuilder.WithSecret(secret)
	return c
}

// WithScopes set the interface to use for fetching user info
func (c *NestedClientsBuilder) WithScope(scopes ...string) *NestedClientsBuilder {
	c.ClientsBuilder.WithScope(scopes...)
	return c
}

// And set the interface to use for fetching user info
func (c *NestedClientsBuilder) And() *ManagerBuilder {
	return c.parent
}

package oauth

import (
	"webapp/core/net/http/security/token"
)

// ManagerBuilder struct to handle basic authentication
type ManagerBuilder struct {
	*Manager
	clientServices []ClientService
	ClientsBuilder *NestedClientsBuilder
}

// NewManagerBuilder Create a new ClientsBuilder
func NewManagerBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		Manager:        &Manager{},
		clientServices: make([]ClientService, 0),
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
	b.clientServices = append(b.clientServices, s)
	return b
}

// WithTokenService set the interface to use for fetch user info
func (b *ManagerBuilder) WithTokenService(ts token.Service) *ManagerBuilder {
	b.tokenService = ts
	return b
}

// WithInMemoryClients set the interface to use for fetching user info
func (b *ManagerBuilder) WithInMemoryClients() *NestedClientsBuilder {
	if b.ClientsBuilder == nil {
		b.ClientsBuilder = newNestedClientsBuilder(b)
	}
	return b.ClientsBuilder
}

// Build Add user service
func (b *ManagerBuilder) Build() *Manager {
	b.clientServices = append([]ClientService{b.ClientsBuilder.Build()}, b.clientServices...)
	b.clientmanager = &ClientManager{
		clientServices: b.clientServices,
	}
	b.controller = NewRestController(NewServiceImpl(b.tokenService, b.clientmanager))
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

// WithScope set the interface to use for fetching user info
func (c *NestedClientsBuilder) WithScope(scopes ...string) *NestedClientsBuilder {
	c.ClientsBuilder.WithScope(scopes...)
	return c
}

// And set the interface to use for fetching user info
func (c *NestedClientsBuilder) And() *ManagerBuilder {
	return c.parent
}

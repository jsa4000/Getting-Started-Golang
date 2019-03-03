package security

import (
	net "webapp/core/net/http"
)

// AuthFacade interface for components to register
type AuthFacade interface {
	Controllers() []net.Controller
}

// Manager returns struct
type Manager struct {
	*Config
	authHandlers   net.Middleware
	filterHandlers net.Middleware
	facades        []AuthFacade
}

// Middleware returns the middleware for the security implementation
func (c *Manager) Middleware() []net.Middleware {
	result := make([]net.Middleware, 0)
	if c.authHandlers != nil {
		result = append(result, c.authHandlers)
	}
	if c.filterHandlers != nil {
		result = append(result, c.filterHandlers)
	}
	return result
}

// Controllers returns the controller for the security implementation
func (c *Manager) Controllers() []net.Controller {
	cc := make([]net.Controller, 0)
	for _, f := range c.facades {
		cc = append(cc, f.Controllers()...)
	}
	return cc
}

// ManagerBuilder returns struct
type ManagerBuilder struct {
	*Manager
	authHndls   []AuthHandler
	filterHndls []FilterHandler
}

// NewBuilder returns new security config
func NewBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		Manager: &Manager{
			Config:  NewConfig(),
			facades: make([]AuthFacade, 0),
		},
		authHndls:   make([]AuthHandler, 0),
		filterHndls: make([]FilterHandler, 0),
	}
}

// WithConfig set middleware to use for security
func (c *ManagerBuilder) WithConfig(config *Config) *ManagerBuilder {
	c.Config = config
	return c
}

// WithAuthentication set middleware to use for authorization
func (c *ManagerBuilder) WithAuthentication(facades ...AuthFacade) *ManagerBuilder {
	c.facades = append(c.facades, facades...)
	return c
}

// WithAuthorization set middleware to use for authorization
func (c *ManagerBuilder) WithAuthorization(method ...AuthHandler) *ManagerBuilder {
	c.authHndls = append(c.authHndls, method...)
	return c
}

// WithFilter set middleware to use for resource filtering
func (c *ManagerBuilder) WithFilter(method ...FilterHandler) *ManagerBuilder {
	c.filterHndls = append(c.filterHndls, method...)
	return c
}

// Build returns manager build
func (c *ManagerBuilder) Build() *Manager {
	if len(c.authHndls) > 0 {
		c.authHandlers = NewMiddleware(c.authHndls, net.PriorityAuth, true)
	}
	if len(c.filterHndls) > 0 {
		c.filterHandlers = NewMiddleware(c.filterHndls, net.PriorityFilters, false)
	}
	return c.Manager
}

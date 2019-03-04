package security

import (
	net "webapp/core/net/http"
)

// Facade interface for components to register
type Facade interface {
	Controllers() []net.Controller
}

// Manager returns struct
type Manager struct {
	*Config
	authHandlers   net.Middleware
	filterHandlers net.Middleware
	facades        []Facade
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

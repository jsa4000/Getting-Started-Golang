package security

import (
	net "webapp/core/net/http"
)

// Manager returns struct
type Manager struct {
	*Config
	handlers     []AuthHandler
	middleware   net.Middleware
	controller   net.Controller
	tokenService TokenService
}

// Middleware returns the middleware for the security implementation
func (c *Manager) Middleware() net.Middleware {
	return c.middleware
}

// Controller returns the controller for the security implementation
func (c *Manager) Controller() net.Controller {
	return c.controller
}

// ManagerBuilder returns struct
type ManagerBuilder struct {
	*Manager
}

// NewBuilder returns new security config
func NewBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		&Manager{
			Config:   NewConfig(),
			handlers: make([]AuthHandler, 0),
		},
	}
}

// WithConfig set middleware to use for security
func (c *ManagerBuilder) WithConfig(config *Config) *ManagerBuilder {
	c.Config = config
	return c
}

// WithHandlers set middleware to use for security
func (c *ManagerBuilder) WithHandlers(method ...AuthHandler) *ManagerBuilder {
	c.handlers = append(c.handlers, method...)
	return c
}

// WithTokenService set the interface to use for fetch user info
func (c *ManagerBuilder) WithTokenService(ts TokenService) *ManagerBuilder {
	c.tokenService = ts
	return c
}

// Build returns manager build
func (c *ManagerBuilder) Build() *Manager {
	c.controller = NewRestController(c.tokenService)
	c.middleware = NewMiddleware(c.handlers)
	return c.Manager
}

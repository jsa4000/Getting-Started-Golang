package security

import (
	net "webapp/core/net/http"
)

// Manager returns struct
type Manager struct {
	*Config
	authentication net.Middleware
	authorization  net.Middleware
	controller     net.Controller
	//tokenService   TokenService
}

// Middleware returns the middleware for the security implementation
func (c *Manager) Middleware() []net.Middleware {
	result := make([]net.Middleware, 0)
	if c.authentication != nil {
		result = append(result, c.authentication)
	}
	if c.authorization != nil {
		result = append(result, c.authorization)
	}
	return result
}

// Controller returns the controller for the security implementation
func (c *Manager) Controller() net.Controller {
	return c.controller
}

// ManagerBuilder returns struct
type ManagerBuilder struct {
	*Manager
	authenticationHndls []AuthHandler
	authorizationHndls  []AuthHandler
}

// NewBuilder returns new security config
func NewBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		Manager: &Manager{
			Config: NewConfig(),
		},
		authenticationHndls: make([]AuthHandler, 0),
		authorizationHndls:  make([]AuthHandler, 0),
	}
}

// WithConfig set middleware to use for security
func (c *ManagerBuilder) WithConfig(config *Config) *ManagerBuilder {
	c.Config = config
	return c
}

// WithAuhtorizationHandlers set middleware to use for security
func (c *ManagerBuilder) WithAuhtorizationHandlers(method ...AuthHandler) *ManagerBuilder {
	c.authenticationHndls = append(c.authenticationHndls, method...)
	return c
}

// WithResourceHandlers set middleware to use for security
func (c *ManagerBuilder) WithResourceHandlers(method ...AuthHandler) *ManagerBuilder {
	c.authorizationHndls = append(c.authorizationHndls, method...)
	return c
}

/*
// WithTokenService set the interface to use for fetch user info
func (c *ManagerBuilder) WithTokenService(ts TokenService) *ManagerBuilder {
	c.tokenService = ts
	return c
}
*/

// Build returns manager build
func (c *ManagerBuilder) Build() *Manager {
	//c.controller = NewRestController(c.tokenService)
	if len(c.authenticationHndls) > 0 {
		c.authentication = NewMiddleware(c.authenticationHndls, net.PriorityAuthentication)
	}
	if len(c.authorizationHndls) > 0 {
		c.authorization = NewMiddleware(c.authorizationHndls, net.PriorityAuthorization)
	}
	return c.Manager
}

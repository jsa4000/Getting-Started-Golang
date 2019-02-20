package security

import (
	net "webapp/core/net/http"
)

// Manager returns struct
type Manager struct {
	middleware  []net.Middleware
	controllers []net.Controller
	config      *Config
}

// New returns new security config
func New(config *Config) net.Security {
	config.Service = NewTokenServiceJwt(config)
	return &Manager{
		middleware: []net.Middleware{
			NewAuthHandlerMiddleware(config),
		},
		controllers: []net.Controller{
			NewRestController(config.Service),
		},
		config: config,
	}
}

// Middleware returns the middleware for the security implementation
func (c *Manager) Middleware() []net.Middleware {
	return c.middleware
}

// Controllers returns the controllers for the security implementation
func (c *Manager) Controllers() []net.Controller {
	return c.controllers
}

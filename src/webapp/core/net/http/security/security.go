package security

import (
	net "webapp/core/net/http"
)

// Config returns struct
type Config struct {
	middleware  []net.Middleware
	controllers []net.Controller
}

// New returns new security config
func New() net.Security {
	return &Config{
		middleware: []net.Middleware{
			NewAuthHandlerMiddleware(),
		},
		controllers: []net.Controller{
			NewRestController(NewServiceImpl()),
		},
	}
}

// Middleware returns the middleware for the security implementation
func (c *Config) Middleware() []net.Middleware {
	return c.middleware
}

// Controllers returns the controllers for the security implementation
func (c *Config) Controllers() []net.Controller {
	return c.controllers
}

package security

import (
	net "webapp/core/net/http"
)

// Config returns struct
type Config struct {
	middleware []net.Middleware
}

// New returns new security config
func New() net.Security {
	return &Config{
		middleware: []net.Middleware{
			NewAuthHandlerMiddleware(),
		},
	}
}

// Middleware returns the middleware for the security implementation
func (c *Config) Middleware() []net.Middleware {
	return c.middleware
}

package security

import (
	"context"
	net "webapp/core/net/http"
)

// UserData structure to identify the user
type UserData struct {
	ID       string
	Name     string
	Email    string
	Password string
	Roles    []string
}

// TokenData data structure for token generation
type TokenData map[string]interface{}

// UserCallback Interface
type UserCallback interface {
	GetUserByName(ctx context.Context, name string) (*UserData, error)
	GetUserByEmail(ctx context.Context, email string) (*UserData, error)
}

// TokenCallback Interface
type TokenCallback interface {
	OnGenerate(ctx context.Context, t *TokenData)
}

// Manager returns struct
type Manager struct {
	middleware  []net.Middleware
	controllers []net.Controller
	config      *Config
}

// New returns new security config
func New(config *Config) net.Security {
	service := NewServiceJwt(config)
	return &Manager{
		middleware: []net.Middleware{
			NewAuthHandlerMiddleware(config, service),
		},
		controllers: []net.Controller{
			NewRestController(service),
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

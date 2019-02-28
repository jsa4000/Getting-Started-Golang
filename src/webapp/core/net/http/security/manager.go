package security

import (
	net "webapp/core/net/http"
	"webapp/core/net/http/security/oauth"
	"webapp/core/net/http/security/token"
)

// Manager returns struct
type Manager struct {
	*Config
	authentication net.Middleware
	authorization  net.Middleware
	controller     net.Controller
	tokenService   token.Service
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
	authorizationHndls  []AuthHandler
	resourceFilterHndls []AuthHandler
}

// NewBuilder returns new security config
func NewBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		Manager: &Manager{
			Config: NewConfig(),
		},
		authorizationHndls:  make([]AuthHandler, 0),
		resourceFilterHndls: make([]AuthHandler, 0),
	}
}

// WithConfig set middleware to use for security
func (c *ManagerBuilder) WithConfig(config *Config) *ManagerBuilder {
	c.Config = config
	return c
}

// WithAuthentication set middleware to use for authorization
func (c *ManagerBuilder) WithAuthentication(authManager *AuthenticationManager) *ManagerBuilder {

	return c
}

// WithAuthorization set middleware to use for authorization
func (c *ManagerBuilder) WithAuthorization(method ...AuthHandler) *ManagerBuilder {
	c.authorizationHndls = append(c.authorizationHndls, method...)
	return c
}

// WithResourceFilter set middleware to use for resource filtering
func (c *ManagerBuilder) WithResourceFilter(method ...AuthHandler) *ManagerBuilder {
	c.resourceFilterHndls = append(c.resourceFilterHndls, method...)
	return c
}

// WithTokenService set the interface to use for fetch user info
func (c *ManagerBuilder) WithTokenService(ts token.Service) *ManagerBuilder {
	c.tokenService = ts
	return c
}

// Build returns manager build
func (c *ManagerBuilder) Build() *Manager {
	c.controller = oauth.NewRestController(c.tokenService)
	if len(c.authorizationHndls) > 0 {
		c.authentication = NewMiddleware(c.authorizationHndls, net.PriorityAuthorization)
	}
	if len(c.resourceFilterHndls) > 0 {
		c.authorization = NewMiddleware(c.resourceFilterHndls, net.PriorityResourceFilters)
	}
	return c.Manager
}

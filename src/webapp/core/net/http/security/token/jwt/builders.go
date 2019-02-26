package jwt

import "webapp/core/net/http/security"

// Builder main app configuration
type Builder struct {
	*AuthHandler
}

// NewBuilder Create a new ServiceImpl
func NewBuilder() *Builder {
	return &Builder{
		&AuthHandler{
			Config:  NewConfig(),
			targets: make([]string, 0),
		},
	}
}

// Withconfig set the interface to use for fetching user info
func (c *Builder) Withconfig(config *Config) *Builder {
	c.Config = config
	return c
}

// WithTargets set the interface to use for fetching user info
func (c *Builder) WithTargets(target ...string) *Builder {
	c.targets = append(c.targets, target...)
	return c
}

// WithUserInfoProvider set the interface to use for fetching user info
func (c *Builder) WithUserInfoProvider(provider security.UserInfoProvider) *Builder {
	c.provider = provider
	return c
}

// Build set User Callback
func (c *Builder) Build() *AuthHandler {
	return c.AuthHandler
}

// ServiceBuilder main app configuration
type ServiceBuilder struct {
	*Service
}

// NewServiceBuilder Create a new ServiceImpl
func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{
		&Service{
			Config: NewConfig(),
		},
	}
}

// Withconfig set the interface to use for fetching user info
func (c *ServiceBuilder) Withconfig(config *Config) *ServiceBuilder {
	c.Config = config
	return c
}

// WithUserInfoProvider set the interface to use for fetching user info
func (c *ServiceBuilder) WithUserInfoProvider(provider security.UserInfoProvider) *ServiceBuilder {
	c.provider = provider
	return c
}

// WithClaimsEnhancer set User Callback
func (c *ServiceBuilder) WithClaimsEnhancer(enchancer ClaimsEnhancer) *ServiceBuilder {
	c.enhancer = enchancer
	return c
}

// Build set User Callback
func (c *ServiceBuilder) Build() *Service {
	return c.Service
}

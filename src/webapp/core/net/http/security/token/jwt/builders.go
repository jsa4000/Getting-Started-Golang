package jwt

import "webapp/core/net/http/security"

// Builder main app configuration
type Builder struct {
	*AuthHandler
	targetsBuilder *NestedTargetsBuilder
}

// NewBuilder Create a new ServiceImpl
func NewBuilder() *Builder {
	return &Builder{
		AuthHandler: &AuthHandler{
			Config: NewConfig(),
		},
	}
}

// NestedTargetsBuilder build struct
type NestedTargetsBuilder struct {
	*security.TargetsBuilder
	parent *Builder
}

// NewTargetsBuilder Create a new TargetsBuilder
func newNestedTargetsBuilder(parent *Builder) *NestedTargetsBuilder {
	return &NestedTargetsBuilder{
		security.NewTargetsBuilder(),
		parent,
	}
}

// Withconfig set the interface to use for fetching user info
func (c *Builder) Withconfig(config *Config) *Builder {
	c.Config = config
	return c
}

// WithTargets set the interface to use for fetching user info
func (c *Builder) WithTargets() *NestedTargetsBuilder {
	c.targetsBuilder = newNestedTargetsBuilder(c)
	return c.targetsBuilder
}

// WithUserInfoService set the interface to use for fetching user info
func (c *Builder) WithUserInfoService(provider security.UserInfoService) *Builder {
	c.provider = provider
	return c
}

// Build set User Callback
func (c *Builder) Build() *AuthHandler {
	c.Targets = c.targetsBuilder.Build()
	return c.AuthHandler
}

// WithURL set the interface to use for fetching user info
func (c *NestedTargetsBuilder) WithURL(url string) *NestedTargetsBuilder {
	c.TargetsBuilder.WithURL(url)
	return c
}

// WithAuthorities set the interface to use for fetching user info
func (c *NestedTargetsBuilder) WithAuthorities(authorities ...string) *NestedTargetsBuilder {
	c.TargetsBuilder.WithAuthorities(authorities...)
	return c
}

// And set the interface to use for fetching user info
func (c *NestedTargetsBuilder) And() *Builder {
	return c.parent
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

// WithUserInfoService set the interface to use for fetching user info
func (c *ServiceBuilder) WithUserInfoService(provider security.UserInfoService) *ServiceBuilder {
	c.provider = provider
	return c
}

// WithClaimsEnhancer set User Callback
func (c *ServiceBuilder) WithClaimsEnhancer(enhancer ClaimsEnhancer) *ServiceBuilder {
	c.enhancer = enhancer
	return c
}

// Build set User Callback
func (c *ServiceBuilder) Build() *Service {
	return c.Service
}

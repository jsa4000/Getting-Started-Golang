package basic

import "webapp/core/net/http/security"

// Builder main app configuration
type Builder struct {
	*AuthHandler
	usersBuilder *NestedDefaultUsersBuilder
}

// NestedDefaultUsersBuilder build struct
type NestedDefaultUsersBuilder struct {
	*DefaultUsersBuilder
	builder *Builder
}

// NewBuilder Create a new ServiceImpl
func NewBuilder() *Builder {
	return &Builder{
		AuthHandler: &AuthHandler{
			Config:  NewConfig(),
			targets: make([]string, 0),
		},
	}
}

// NewDefaultUsersBuilder Create a new DefaultUsersBuilder
func newNestedDefaultUsersBuilder(builder *Builder) *NestedDefaultUsersBuilder {
	return &NestedDefaultUsersBuilder{
		NewDefaultUsersBuilder(),
		builder,
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

// WithLocalUsers set the interface to use for fetching user info
func (c *Builder) WithLocalUsers() *NestedDefaultUsersBuilder {
	c.usersBuilder = newNestedDefaultUsersBuilder(c)
	return c.usersBuilder
}

// WithUserInfoProvider set the interface to use for fetching user info
func (c *Builder) WithUserInfoProvider(provider security.UserInfoProvider) *Builder {
	c.provider = provider
	return c
}

// Build set User Callback
func (c *Builder) Build() *AuthHandler {
	if len(c.usersBuilder.Users) > 0 {
		c.local = c.usersBuilder.DefaultUsers
	}
	return c.AuthHandler
}

// WithUser set the interface to use for fetching user info
func (c *NestedDefaultUsersBuilder) WithUser(name string) *NestedDefaultUsersBuilder {
	c.DefaultUsersBuilder.WithUser(name)
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *NestedDefaultUsersBuilder) WithPassword(password string) *NestedDefaultUsersBuilder {
	c.DefaultUsersBuilder.WithPassword(password)
	return c
}

// WithRoles set the interface to use for fetching user info
func (c *NestedDefaultUsersBuilder) WithRoles(roles []string) *NestedDefaultUsersBuilder {
	c.DefaultUsersBuilder.WithRoles(roles)
	return c
}

// And set the interface to use for fetching user info
func (c *NestedDefaultUsersBuilder) And() *Builder {
	c.Build()
	return c.builder
}

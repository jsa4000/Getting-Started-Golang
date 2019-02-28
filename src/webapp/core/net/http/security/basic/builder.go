package basic

import "webapp/core/net/http/security"

// Builder main app configuration
type Builder struct {
	*AuthHandler
	usersBuilder *NestedAuthorizedUsersBuilder
}

// NestedAuthorizedUsersBuilder build struct
type NestedAuthorizedUsersBuilder struct {
	*security.AuthorizedUsersBuilder
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

// NewAuthorizedUsersBuilder Create a new AuthorizedUsersBuilder
func newNestedAuthorizedUsersBuilder(builder *Builder) *NestedAuthorizedUsersBuilder {
	return &NestedAuthorizedUsersBuilder{
		security.NewAuthorizedUsersBuilder(),
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
func (c *Builder) WithLocalUsers() *NestedAuthorizedUsersBuilder {
	c.usersBuilder = newNestedAuthorizedUsersBuilder(c)
	return c.usersBuilder
}

// WithUserInfoService set the interface to use for fetching user info
func (c *Builder) WithUserInfoService(provider security.UserInfoService) *Builder {
	c.provider = provider
	return c
}

// Build set User Callback
func (c *Builder) Build() *AuthHandler {
	if len(c.usersBuilder.Users) > 0 {
		c.local = c.usersBuilder.AuthorizedUsers
	}
	return c.AuthHandler
}

// WithUser set the interface to use for fetching user info
func (c *NestedAuthorizedUsersBuilder) WithUser(name string) *NestedAuthorizedUsersBuilder {
	c.AuthorizedUsersBuilder.WithUser(name)
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *NestedAuthorizedUsersBuilder) WithPassword(password string) *NestedAuthorizedUsersBuilder {
	c.AuthorizedUsersBuilder.WithPassword(password)
	return c
}

// WithRoles set the interface to use for fetching user info
func (c *NestedAuthorizedUsersBuilder) WithRoles(roles []string) *NestedAuthorizedUsersBuilder {
	c.AuthorizedUsersBuilder.WithRoles(roles)
	return c
}

// And set the interface to use for fetching user info
func (c *NestedAuthorizedUsersBuilder) And() *Builder {
	c.Build()
	return c.builder
}

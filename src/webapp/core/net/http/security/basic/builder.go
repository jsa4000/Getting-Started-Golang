package basic

import "webapp/core/net/http/security"

// Builder main app configuration
type Builder struct {
	*AuthHandler
	usersBuilder *LocalUserBuilder
}

// LocalUserBuilder main app configuration
type LocalUserBuilder struct {
	builder *Builder
	users   *LocalUsers
	current *security.UserInfo
}

// NewBuilder Create a new ServiceImpl
func newLocalUserBuilder(builder *Builder) *LocalUserBuilder {
	return &LocalUserBuilder{
		builder: builder,
		users:   NewlocalUsers(),
	}
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
func (c *Builder) WithLocalUsers() *LocalUserBuilder {
	c.usersBuilder = newLocalUserBuilder(c)
	return c.usersBuilder
}

// WithUserInfoProvider set the interface to use for fetching user info
func (c *Builder) WithUserInfoProvider(provider security.UserInfoProvider) *Builder {
	c.provider = provider
	return c
}

// Build set User Callback
func (c *Builder) Build() *AuthHandler {
	if len(c.usersBuilder.users.Users) > 0 {
		c.local = c.usersBuilder.users
	}
	return c.AuthHandler
}

// WithUser set the interface to use for fetching user info
func (c *LocalUserBuilder) WithUser(name string) *LocalUserBuilder {
	if c.current != nil {
		c.users.Users[c.current.Name] = c.current
	}
	c.current = &security.UserInfo{
		Name: name,
	}
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *LocalUserBuilder) WithPassword(password string) *LocalUserBuilder {
	c.current.Password = password
	return c
}

// And set the interface to use for fetching user info
func (c *LocalUserBuilder) And() *Builder {
	if c.current != nil {
		c.users.Users[c.current.Name] = c.current
	}
	return c.builder
}

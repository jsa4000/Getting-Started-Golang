package basic

import "webapp/core/net/http/security"

// Builder main app configuration
type Builder struct {
	*AuthHandler
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

// WithUserInfoService set the interface to use for fetching user info
func (c *Builder) WithUserInfoService(s security.UserInfoService) *Builder {
	c.service = s
	return c
}

// Build set User Callback
func (c *Builder) Build() *AuthHandler {
	return c.AuthHandler
}

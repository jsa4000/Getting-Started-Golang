package oauth

import "webapp/core/net/http/security"

// Builder main app configuration
type Builder struct {
	*Service
}

// NewBuilder Create a new ServiceImpl
func NewBuilder() *Builder {
	return &Builder{
		&Service{
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

// WithUserFetcher set the interface to use for fetching user info
func (c *Builder) WithUserFetcher(fecther security.UserFetcher) *Builder {
	c.userFetcher = fecther
	return c
}

// Build set User Callback
func (c *Builder) Build() *Service {
	return c.Service
}

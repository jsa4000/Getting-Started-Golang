package roles

import "webapp/core/net/http/security"

// Builder main app configuration
type Builder struct {
	*Handler
	targetsBuilder *NestedTargetsBuilder
}

// NewBuilder Create a new ServiceImpl
func NewBuilder() *Builder {
	return &Builder{
		Handler: &Handler{
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

// WithConfig set the interface to use for fetching user info
func (c *Builder) WithConfig(config *Config) *Builder {
	c.Config = config
	return c
}

// WithPriority set the interface to use for fetching user info
func (c *Builder) WithPriority(priority int) *Builder {
	c.Prior = priority
	return c
}

// WithTargets set the interface to use for fetching user info
func (c *Builder) WithTargets() *NestedTargetsBuilder {
	c.targetsBuilder = newNestedTargetsBuilder(c)
	return c.targetsBuilder
}

// Build set User Callback
func (c *Builder) Build() *Handler {
	c.Targets = c.targetsBuilder.Build()
	return c.Handler
}

// WithURL set the interface to use for fetching user info
func (c *NestedTargetsBuilder) WithURL(url string) *NestedTargetsBuilder {
	c.TargetsBuilder.WithURL(url)
	return c
}

// WithAuthority set the interface to use for fetching user info
func (c *NestedTargetsBuilder) WithAuthority(authorities ...string) *NestedTargetsBuilder {
	c.TargetsBuilder.WithAuthority(authorities...)
	return c
}

// And set the interface to use for fetching user info
func (c *NestedTargetsBuilder) And() *Builder {
	return c.parent
}

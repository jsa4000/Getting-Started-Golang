package access

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
	*TargetsBuilder
	parent *Builder
}

// NewTargetsBuilder Create a new TargetsBuilder
func newNestedTargetsBuilder(parent *Builder) *NestedTargetsBuilder {
	return &NestedTargetsBuilder{
		NewTargetsBuilder(),
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

// WithOrigin set the interface to use for fetching user info
func (c *NestedTargetsBuilder) WithOrigin(origin string) *NestedTargetsBuilder {
	c.TargetsBuilder.WithOrigin(origin)
	return c
}

// WithCredentials set the interface to use for fetching user info
func (c *NestedTargetsBuilder) WithCredentials(enable bool) *NestedTargetsBuilder {
	c.TargetsBuilder.WithCredentials(enable)
	return c
}

// WithMethods set the interface to use for fetching user info
func (c *NestedTargetsBuilder) WithMethods(methods ...string) *NestedTargetsBuilder {
	c.TargetsBuilder.WithMethods(methods...)
	return c
}

// WithHeaders set the interface to use for fetching user info
func (c *NestedTargetsBuilder) WithHeaders(headers ...string) *NestedTargetsBuilder {
	c.TargetsBuilder.WithHeaders(headers...)
	return c
}

// Allow set the interface to use for fetching user info
func (c *NestedTargetsBuilder) Allow() *NestedTargetsBuilder {
	c.TargetsBuilder.Allow()
	return c
}

// And set the interface to use for fetching user info
func (c *NestedTargetsBuilder) And() *Builder {
	return c.parent
}

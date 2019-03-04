package access

import "webapp/core/net/http/security"

// Target structure to retrieve (fetch) the Target information
type Target struct {
	security.TargetBase
	Allow        bool
	Origin       string
	Crendentials int
	Methods      []string
	Headers      []string
}

func newTarget(url string) *Target {
	return &Target{
		TargetBase: security.TargetBase{
			URL:         url,
			Authorities: make([]string, 0),
		},
		Allow:        false,
		Origin:       AllowAnyOrigin,
		Crendentials: -1,
		Methods:      make([]string, 0),
		Headers:      make([]string, 0),
	}
}

// TargetsBuilder build struct
type TargetsBuilder struct {
	security.Targets
	current *Target
}

// NewTargetsBuilder Create a new DefaultUsersBuilder
func NewTargetsBuilder() *TargetsBuilder {
	return &TargetsBuilder{
		Targets: make([]security.Target, 0),
	}
}

// WithURL set the interface to use for fetching user info
func (c *TargetsBuilder) WithURL(url string) *TargetsBuilder {
	if c.current != nil {
		c.Targets = append(c.Targets, c.current)
	}
	c.current = newTarget(url)
	return c
}

// WithOrigin set the interface to use for fetching user info
func (c *TargetsBuilder) WithOrigin(origin string) *TargetsBuilder {
	c.current.Origin = origin
	return c
}

// WithCredentials set the interface to use for fetching user info
func (c *TargetsBuilder) WithCredentials(enable bool) *TargetsBuilder {
	c.current.Crendentials = toInt(enable)
	return c
}

// WithMethods set the interface to use for fetching user info
func (c *TargetsBuilder) WithMethods(methods ...string) *TargetsBuilder {
	c.current.Methods = append(c.current.Methods, methods...)
	return c
}

// WithHeaders set the interface to use for fetching user info
func (c *TargetsBuilder) WithHeaders(headers ...string) *TargetsBuilder {
	c.current.Headers = append(c.current.Headers, headers...)
	return c
}

// Allow set the interface to use for fetching user info
func (c *TargetsBuilder) Allow() *TargetsBuilder {
	c.current.Allow = true
	return c
}

// Build default users struct
func (c *TargetsBuilder) Build() security.Targets {
	if c.current != nil {
		c.Targets = append(c.Targets, c.current)
	}
	return c.Targets
}

func toInt(enabled bool) int {
	if enabled {
		return 1
	}
	return 0
}

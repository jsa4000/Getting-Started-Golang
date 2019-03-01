package security

// Target structure to retrieve (fetch) the Target information
type Target struct {
	URL         string   `json:"url"`
	Authorities []string `json:"authorities"`
}

// Targets to implement the UserinfoProvider
type Targets []*Target

// TargetsBuilder build struct
type TargetsBuilder struct {
	Targets
	current *Target
}

// NewTargetsBuilder Create a new DefaultUsersBuilder
func NewTargetsBuilder() *TargetsBuilder {
	return &TargetsBuilder{
		Targets: make([]*Target, 0),
	}
}

// Matches any target with the given URL
func (c *Targets) Matches(url string) (*Target, bool) {

	return nil, false
}

// WithURL set the interface to use for fetching user info
func (c *TargetsBuilder) WithURL(url string) *TargetsBuilder {
	if c.current != nil {
		c.Targets = append(c.Targets, c.current)
	}
	c.current = &Target{
		URL: url,
	}
	return c
}

// WithAuthorities set the interface to use for fetching user info
func (c *TargetsBuilder) WithAuthorities(authorities []string) *TargetsBuilder {
	c.current.Authorities = authorities
	return c
}

// Build default users struct
func (c *TargetsBuilder) Build() *Targets {
	if c.current != nil {
		c.Targets = append(c.Targets, c.current)
	}
	return &c.Targets
}

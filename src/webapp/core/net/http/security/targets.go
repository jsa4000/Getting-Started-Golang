package security

import (
	net "webapp/core/net/http"
)

// Matcher Interface
type Matcher interface {
	Matches(url string) (*Target, bool)
}

// Target structure to retrieve (fetch) the Target information
type Target struct {
	URL         string   `json:"url"`
	Authorities []string `json:"authorities"`
}

// Any any target with the given URL
func (t *Target) Any(authorities []string) bool {
	if len(t.Authorities) == 0 {
		return true
	}
	if len(authorities) == 0 {
		return false
	}
	for _, ta := range t.Authorities {
		for _, a := range authorities {
			if ta == a {
				return true
			}
		}
	}
	return false
}

func newTarget(url string) *Target {
	return &Target{
		URL:         url,
		Authorities: make([]string, 0),
	}
}

// Targets to implement the Matcher interface
type Targets []*Target

// Matches any target with the given URL
func (c *Targets) Matches(url string) (*Target, bool) {
	for _, t := range *c {
		if net.MatchesURL(url, t.URL) {
			return t, true
		}
	}
	return nil, false
}

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

// WithURL set the interface to use for fetching user info
func (c *TargetsBuilder) WithURL(url string) *TargetsBuilder {
	if c.current != nil {
		c.Targets = append(c.Targets, c.current)
	}
	c.current = newTarget(url)
	return c
}

// WithAuthorities set the interface to use for fetching user info
func (c *TargetsBuilder) WithAuthorities(authorities ...string) *TargetsBuilder {
	c.current.Authorities = append(c.current.Authorities, authorities...)
	return c
}

// Build default users struct
func (c *TargetsBuilder) Build() *Targets {
	if c.current != nil {
		c.Targets = append(c.Targets, c.current)
	}
	return &c.Targets
}

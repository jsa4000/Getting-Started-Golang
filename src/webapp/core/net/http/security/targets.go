package security

import (
	net "webapp/core/net/http"
)

// Target interface to retrieve URL
type Target interface {
	Matches(url string) bool
	Any(authorities []string) bool
}

// Target structure to retrieve (fetch) the Target information
type target struct {
	url         string
	authorities []string
}

func newTarget(url string) *target {
	return &target{
		url:         url,
		authorities: make([]string, 0),
	}
}

// Matches any target with the given URL
func (t *target) Matches(url string) bool {
	if net.MatchesURL(url, t.url) {
		return true
	}
	return false
}

// Any any target with the given URL
func (t *target) Any(authorities []string) bool {
	if len(t.authorities) == 0 {
		return true
	}
	if len(authorities) == 0 {
		return false
	}
	for _, ta := range t.authorities {
		for _, a := range authorities {
			if ta == a {
				return true
			}
		}
	}
	return false
}

// Targets to implement the Matcher interface
type Targets []Target

// TargetsBuilder build struct
type TargetsBuilder struct {
	Targets
	current *target
}

// NewTargetsBuilder Create a new DefaultUsersBuilder
func NewTargetsBuilder() *TargetsBuilder {
	return &TargetsBuilder{
		Targets: make([]Target, 0),
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

// WithAuthority set the interface to use for fetching user info
func (c *TargetsBuilder) WithAuthority(authorities ...string) *TargetsBuilder {
	c.current.authorities = append(c.current.authorities, authorities...)
	return c
}

// Build default users struct
func (c *TargetsBuilder) Build() Targets {
	if c.current != nil {
		c.Targets = append(c.Targets, c.current)
	}
	return c.Targets
}

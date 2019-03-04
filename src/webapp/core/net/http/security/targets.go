package security

import (
	"fmt"
	net "webapp/core/net/http"
)

// Target interface to retrieve URL
type Target interface {
	Matches(url string) bool
	Any(authorities []string) bool
}

// TargetBase structure to retrieve (fetch) the Target information
type TargetBase struct {
	URL         string
	Authorities []string
}

// NewTargetBase Creates new Target Base
func NewTargetBase(url string) *TargetBase {
	return &TargetBase{
		URL:         url,
		Authorities: make([]string, 0),
	}
}

// Matches any target with the given URL
func (t *TargetBase) Matches(url string) bool {
	if net.MatchesURL(url, t.URL) {
		return true
	}
	return false
}

// Any any target with the given URL
func (t *TargetBase) Any(authorities []string) bool {
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

// Targets to implement the Matcher interface
type Targets []Target

func (t Targets) String() string {
	result := ""
	for _, item := range t {
		if target, ok := item.(*TargetBase); ok {
			result += fmt.Sprintf("%v", target)
		}
	}
	return result
}

// TargetsBuilder build struct
type TargetsBuilder struct {
	Targets
	current *TargetBase
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
	c.current = NewTargetBase(url)
	return c
}

// WithAuthority set the interface to use for fetching user info
func (c *TargetsBuilder) WithAuthority(authorities ...string) *TargetsBuilder {
	c.current.Authorities = append(c.current.Authorities, authorities...)
	return c
}

// Build default users struct
func (c *TargetsBuilder) Build() Targets {
	if c.current != nil {
		c.Targets = append(c.Targets, c.current)
	}
	return c.Targets
}

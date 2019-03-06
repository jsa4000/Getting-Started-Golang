package security

import (
	net "webapp/core/net/http"
)

// ManagerBuilder returns struct
type ManagerBuilder struct {
	*Manager
	authHndls   []Handler
	filterHndls []FilterHandler
}

// NewBuilder returns new security config
func NewBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		Manager: &Manager{
			Config:  NewConfig(),
			facades: make([]Facade, 0),
		},
		authHndls:   make([]Handler, 0),
		filterHndls: make([]FilterHandler, 0),
	}
}

// WithConfig set middleware to use for security
func (c *ManagerBuilder) WithConfig(config *Config) *ManagerBuilder {
	c.Config = config
	return c
}

// WithAuthentication set middleware to use for authorization
func (c *ManagerBuilder) WithAuthentication(facades ...Facade) *ManagerBuilder {
	c.facades = append(c.facades, facades...)
	return c
}

// WithAuthorization set middleware to use for authorization
func (c *ManagerBuilder) WithAuthorization(method ...Handler) *ManagerBuilder {
	c.authHndls = append(c.authHndls, method...)
	return c
}

// WithFilter set middleware to use for resource filtering
func (c *ManagerBuilder) WithFilter(method ...FilterHandler) *ManagerBuilder {
	c.filterHndls = append(c.filterHndls, method...)
	return c
}

// Build returns manager build
func (c *ManagerBuilder) Build() *Manager {
	if len(c.authHndls) > 0 {
		c.authHandlers = NewMiddleware(SortFilters(c.authHndls, true), net.PriorityAuth, true)
	}
	if len(c.filterHndls) > 0 {
		c.filterHandlers = NewMiddleware(SortFilters(c.filterHndls, true), net.PriorityFilters, false)
	}
	return c.Manager
}

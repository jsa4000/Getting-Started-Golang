package mngmt

import (
	net "webapp/core/net/http"
)

const (
	// DefaultRootPath root for the API to expose the endpoints
	DefaultRootPath = "/mngmt"
	// DefaultRefreshTime default refresh time to take snapshots
	DefaultRefreshTime = 5 // seconds
)

// Manager struct to handle basic authentication
type Manager struct {
	controller  net.Controller
	path        string
	metrics     bool
	runtime     bool
	health      bool
	refreshTime int
	service     Service
}

// Controller Returns the controller
func (m *Manager) Controller() net.Controller {
	return m.controller
}

// ManagerBuilder struct to handle basic authentication
type ManagerBuilder struct {
	*Manager
}

// NewManagerBuilder Create a new ClientsBuilder
func NewManagerBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		Manager: &Manager{
			path:        DefaultRootPath,
			metrics:     false,
			runtime:     true,
			health:      true,
			refreshTime: DefaultRefreshTime,
		},
	}
}

// WithHealth changes the root path
func (b *ManagerBuilder) WithHealth(enabled bool) *ManagerBuilder {
	b.health = enabled
	return b
}

// WithRuntime enables the metrics
func (b *ManagerBuilder) WithRuntime(enabled bool) *ManagerBuilder {
	b.runtime = enabled
	return b
}

// WithMetrics enables the metrics
func (b *ManagerBuilder) WithMetrics(enabled bool) *ManagerBuilder {
	b.metrics = enabled
	return b
}

// WithRootPath changes the root path
func (b *ManagerBuilder) WithRootPath(path string) *ManagerBuilder {
	b.path = path
	return b
}

// WithRefreshTime changes the root path
func (b *ManagerBuilder) WithRefreshTime(seconds int) *ManagerBuilder {
	b.refreshTime = seconds
	return b
}

// Build Add user service
func (b *ManagerBuilder) Build() *Manager {
	service := NewServiceImpl(b.health, b.metrics, b.runtime, b.refreshTime)
	b.controller = NewRestController(service, b.path)
	return b.Manager
}

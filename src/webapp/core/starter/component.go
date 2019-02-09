package starter

import (
	"context"
)

// Component interface for components to register
type Component interface {
	Init(ctx context.Context)
	Close(ctx context.Context)
}

// Init initialize defaults values
func Init(ctx context.Context) {
	// Initialize all the components autoloaded
	for _, c := range Components() {
		c.Init(ctx)
	}
}

// Shutdown initialize defaults values
func Shutdown(ctx context.Context) {
	// Shutdown all the components autoloaded
	for _, c := range Components() {
		c.Close(ctx)
	}
}

package starters

import (
	"context"
	"fmt"
)

// Component interface for components to register
type Component interface {
	Init(ctx context.Context)
	Close(ctx context.Context)
}

var control = NewControl()

// NewControl creates a new instance
func NewControl() *Control {
	return &Control{
		Components: make(map[string]Component),
	}
}

// Control struct to register the components
type Control struct {
	Components map[string]Component
}

// Register current component
func Register(key string, comp Component) error {
	return control.Register(key, comp)
}

// Clean removes all the components registered
func Clean() {
	control.Clean()
}

// Register current component
func (c *Control) Register(key string, comp Component) error {
	_, ok := c.Components[key]
	if ok {
		return fmt.Errorf("Component %T already registered", comp)
	}
	c.Components[key] = comp
	return nil
}

// Clean removes all the components registered
func (c *Control) Clean() {
	c.Components = make(map[string]Component)
}

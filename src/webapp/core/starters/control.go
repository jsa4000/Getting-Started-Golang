package starters

import (
	"fmt"
)

var control = NewControl()

// Control struct to register the components
type Control struct {
	components map[string]Component
}

// NewControl creates a new instance
func NewControl() *Control {
	return &Control{
		components: make(map[string]Component),
	}
}

// Components Get current components registered
func Components() map[string]Component {
	return control.components
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
	_, ok := c.components[key]
	if ok {
		return fmt.Errorf("Component %T already registered", comp)
	}
	c.components[key] = comp
	return nil
}

// Clean removes all the components registered
func (c *Control) Clean() {
	c.components = make(map[string]Component)
}

package starter

import (
	"context"
	"webapp/core/starter"
	"webapp/core/validation"
	"webapp/core/validation/goplayground"
)

// Wrapper Global Mongo wrapper
var component = New()

// Component for mongo
type Component struct {
	router *goplayground.Validator
}

// New creates a new component to register the wrapped
func New() *Component {
	result := &Component{
		router: goplayground.New(),
	}
	validation.SetGlobal(result.router)
	starter.Register("validator", result)
	return result
}

// Init function that will be called after register the component
func (c *Component) Init(_ context.Context) {

}

// Close function that will be called at the end of the application
func (c *Component) Close(_ context.Context) {

}

// Priority Get the priority to be initialized
func (c *Component) Priority() int {
	return starter.PriorityValidation
}

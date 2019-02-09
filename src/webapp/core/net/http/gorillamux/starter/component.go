package starter

import (
	"context"
	"webapp/core/net/http"
	router "webapp/core/net/http/gorillamux"
	"webapp/core/starter"
)

// Wrapper Global Mongo wrapper
var component = New()

// Component for mongo
type Component struct {
	router *router.Router
}

// New creates a new component to register the wrapped
func New() *Component {
	result := &Component{
		router: router.New(),
	}
	http.SetGlobal(result.router)
	starter.Register("router", result)
	return result
}

// Init function that will be called after register the component
func (c *Component) Init(_ context.Context) {

}

// Close function that will be called at the end of the application
func (c *Component) Close(_ context.Context) {

}

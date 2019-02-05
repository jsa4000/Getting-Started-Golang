package starter

import (
	"context"
	"webapp/core/net/http"
	router "webapp/core/net/http/httprouter"
	"webapp/core/starters"
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
	starters.Register("router", result)
	return result
}

// Init function that will be called after register the component
func (c *Component) Init(_ context.Context) {

}

// Close fucntion that willbe called at the end of the application
func (c *Component) Close(_ context.Context) {

}

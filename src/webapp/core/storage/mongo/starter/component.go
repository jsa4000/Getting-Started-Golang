package starter

import (
	"context"
	"webapp/core/starter"
	"webapp/core/storage/mongo"
)

// Wrapper Global Mongo wrapper
var component = New()

// Component for mongo
type Component struct {
	wrapper *mongo.Wrapper
}

// New creates a new component to register the wrapped
func New() *Component {
	result := &Component{
		wrapper: mongo.New(),
	}
	mongo.SetGlobal(result.wrapper)
	starter.Register("mongo", result)
	return result
}

// Init function that will be called after register the component
func (c *Component) Init(ctx context.Context) {
	config := mongo.LoadConfig()
	c.wrapper.Connect(ctx, config.URL)
}

// Close function that will be called at the end of the application
func (c *Component) Close(ctx context.Context) {
	c.wrapper.Disconnect(ctx)
}

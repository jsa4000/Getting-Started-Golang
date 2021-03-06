package starter

import (
	"context"
	"webapp/core/mngmt"
	"webapp/core/starter"
	"webapp/core/store/cache"
	"webapp/core/store/redis"
)

// Wrapper Global Mongo wrapper
var component = New()

// Component for ewdis
type Component struct {
	wrapper *redis.Wrapper
}

// New creates a new component to register the wrapped
func New() *Component {
	result := &Component{
		wrapper: redis.New(),
	}
	cache.SetGlobal(result.wrapper)
	redis.SetGlobal(result.wrapper)
	starter.Register("redis", result)
	return result
}

// Init function that will be called after register the component
func (c *Component) Init(ctx context.Context) {
	config := redis.LoadConfig()
	c.wrapper.Connect(ctx, config)
}

// Close function that will be called at the end of the application
func (c *Component) Close(ctx context.Context) {
	c.wrapper.Disconnect(ctx)
}

// Priority Get the priority to be initialized
func (c *Component) Priority() int {
	return starter.PriorityStorage
}

// Status Load config from file
func (c *Component) Status() *mngmt.Health {
	status := mngmt.StatusOk
	if !c.wrapper.Connected() {
		status = mngmt.StatusError
	}
	return &mngmt.Health{
		Name:   "Redis",
		Status: status,
	}
}

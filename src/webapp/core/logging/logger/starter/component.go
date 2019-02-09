package starter

import (
	"context"
	"webapp/core/logging"
	"webapp/core/logging/logger"
	"webapp/core/starter"
)

// Wrapper Global Mongo wrapper
var component = New()

// Component for mongo
type Component struct {
	logger *logger.Logger
}

// New creates a new component to register the component
func New() *Component {
	result := &Component{
		logger: logger.New(),
	}
	logging.SetGlobal(result.logger)
	starter.Register("logger", result)
	return result
}

// Init function that will be called after register the component
func (c *Component) Init(_ context.Context) {
	c.logger.SetLevel(logging.DebugLevel)
	c.logger.SetFormatter(logging.TextFormat)
}

// Close fucntion that willbe called at the end of the application
func (c *Component) Close(_ context.Context) {

}

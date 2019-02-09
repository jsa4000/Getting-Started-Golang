package starter

import (
	"context"
	"webapp/core/logging"
	"webapp/core/logging/logrus"
	"webapp/core/starter"
)

// Wrapper Global Mongo wrapper
var component = New()

// Component for mongo
type Component struct {
	logger *logrus.Logger
}

// New creates a new component to register the component
func New() *Component {
	result := &Component{
		logger: logrus.New(),
	}
	logging.SetGlobal(result.logger)
	starter.Register("logger", result)
	return result
}

// Init function that will be called after register the component
func (c *Component) Init(_ context.Context) {
	config := logging.LoadConfig()
	c.logger.SetLevel(config.Level)
	c.logger.SetFormatter(config.Format)
	c.logger.SetOutput(config.Output)
}

// Close function that will called at the end of the application
func (c *Component) Close(_ context.Context) {

}

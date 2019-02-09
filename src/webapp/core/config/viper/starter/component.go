package starter

import (
	"context"
	"log"
	"webapp/core/config"
	"webapp/core/config/viper"
	"webapp/core/starter"
)

// Wrapper Global Mongo wrapper
var component = New()

// Component for mongo
type Component struct {
	parser *viper.Parser
}

// New creates a new component to register the component
func New() *Component {
	result := &Component{
		parser: viper.New(),
	}
	config.SetGlobal(result.parser)
	starter.Register("config", result)
	return result
}

// Init function that will be called after register the component
func (c *Component) Init(_ context.Context) {
	err := c.parser.LoadFromFile("webapp.yaml", ".")
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

// Close fucntion that willbe called at the end of the application
func (c *Component) Close(_ context.Context) {

}

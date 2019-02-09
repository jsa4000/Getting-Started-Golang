package starter

import (
	"webapp/core/config"
)

// Config config for configuration
type Config struct {
	Name string `config:"app.name:MyApp"`
}

// LoadConfig Load config from file
func LoadConfig() *Config {
	c := Config{}
	config.ReadFields(&c)
	return &c
}

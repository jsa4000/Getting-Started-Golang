package security

import (
	"webapp/core/config"
)

// Config main app configuration
type Config struct {
	Enabled string `config:"security.enabled:true"`
}

// NewConfig creates new config
func NewConfig() *Config {
	return loadConfig()
}

// LoadConfig Load config from file
func loadConfig() *Config {
	c := Config{}
	config.ReadFields(&c)
	return &c
}

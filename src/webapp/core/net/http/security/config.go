package security

import (
	"webapp/core/config"
)

// Config main app configuration
type Config struct {
	expiretime int `security:"security.expiretime:60000"`
}

// LoadConfig Load config from file
func LoadConfig() *Config {
	c := Config{}
	config.ReadFields(&c)
	return &c
}

package oauth

import (
	"webapp/core/config"
)

// Config main app configuration
type Config struct {
	ClientID     string `config:"security.clientid:trusted-client"`
	ClientSecret string `config:"security.clientsecret:mypassword$"`
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

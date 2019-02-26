package jwt

import "webapp/core/config"

// Config main app configuration
type Config struct {
	Issuer         string `config:"security.issuer:webapp-oauth"`
	ExpirationTime int    `config:"security.expirationtime:60000"`
	SecretKey      string `config:"security.secretkey:mypassword$"`
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

package security

import (
	"webapp/core/config"
)

// Config main app configuration
type Config struct {
	expiretime int `security:"security.expiretime:60000"`
	uc         UserCallback
	tc         TokenCallback
}

// NewConfig creates new config
func NewConfig() *Config {
	return LoadConfig()
}

// LoadConfig Load config from file
func LoadConfig() *Config {
	c := Config{}
	config.ReadFields(&c)
	return &c
}

// ExpireTime set User Callback
func (c *Config) ExpireTime(t int) *Config {
	c.expiretime = t
	return c
}

// WithUserCallback set User Callback
func (c *Config) WithUserCallback(uc UserCallback) *Config {
	c.uc = uc
	return c
}

// WithTokenCallback set User Callback
func (c *Config) WithTokenCallback(tc TokenCallback) *Config {
	c.tc = tc
	return c
}

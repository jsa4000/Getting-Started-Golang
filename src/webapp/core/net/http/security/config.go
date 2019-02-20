package security

import (
	"webapp/core/config"
)

// Config main app configuration
type Config struct {
	Issuer         string `config:"security.issuer:webapp-oauth"`
	ExpirationTime int    `config:"security.expirationtime:60000"`
	SecretKey      string `config:"security.secretkey:mypassword$"`
	ClientID       string `config:"security.clientid:trusted-client"`
	ClientSecret   string `config:"security.clientsecret:mypassword$"`
	UserFetcher    UserFetcher
	TokenEnhancer  TokenEnhancer
}

// ConfigBuilder main app configuration
type ConfigBuilder struct {
	*Config
}

// NewConfigBuilder creates new config
func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		loadConfig(),
	}
}

// LoadConfig Load config from file
func loadConfig() *Config {
	c := Config{}
	config.ReadFields(&c)
	return &c
}

// WithIssuer set User Callback
func (c *ConfigBuilder) WithIssuer(val string) *ConfigBuilder {
	c.Issuer = val
	return c
}

// WithExpirationTime set User Callback
func (c *ConfigBuilder) WithExpirationTime(t int) *ConfigBuilder {
	c.ExpirationTime = t
	return c
}

// WithSecretKey set User Callback
func (c *ConfigBuilder) WithSecretKey(val string) *ConfigBuilder {
	c.SecretKey = val
	return c
}

// WithClientID set User Callback
func (c *ConfigBuilder) WithClientID(val string) *ConfigBuilder {
	c.ClientID = val
	return c
}

// WithClientSecret set User Callback
func (c *ConfigBuilder) WithClientSecret(val string) *ConfigBuilder {
	c.ClientSecret = val
	return c
}

// WithUserFetcher set User Callback
func (c *ConfigBuilder) WithUserFetcher(fecther UserFetcher) *ConfigBuilder {
	c.UserFetcher = fecther
	return c
}

// WithTokenEnhancer set User Callback
func (c *ConfigBuilder) WithTokenEnhancer(enchancer TokenEnhancer) *ConfigBuilder {
	c.TokenEnhancer = enchancer
	return c
}

// Build set User Callback
func (c *ConfigBuilder) Build() *Config {
	return c.Config
}

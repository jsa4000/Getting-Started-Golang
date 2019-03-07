package redis

import (
	"webapp/core/config"
)

// Config main app configuration
type Config struct {
	URL        string `config:"cache.redis.url:dockerhost:6379"`
	Password   string `config:"cache.redis.password:root"`
	Database   int    `config:"cache.redis.databse:0"`
	MaxRetries int    `config:"cache.redis.maxretries:0"`
}

// LoadConfig Load config from file
func LoadConfig() *Config {
	c := Config{}
	config.ReadFields(&c)
	return &c
}

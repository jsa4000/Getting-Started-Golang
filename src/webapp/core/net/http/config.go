package http

import (
	"time"
	"webapp/core/config"
)

// Config main app configuration
type Config struct {
	Port         int           `config:"http.port:8080"`
	WriteTimeout int           `config:"http.writeTimeout:60"`
	ReadTimeout  int           `config:"http.readTimeout:60"`
	IdleTimeout  time.Duration `config:"http.idleTimeout:60"`
}

// LoadConfig Load config from file
func LoadConfig() *Config {
	c := Config{}
	config.ReadFields(&c)
	return &c
}

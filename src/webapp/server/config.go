package server

import (
	"reflect"
	"time"
	"webapp/config"
)

// Config main app configuration
type Config struct {
	Name         string        `config:"app.name:ServerApp"`
	LogLevel     string        `config:"logging.level:info"`
	Port         int           `config:"server.port:8080"`
	WriteTimeout int           `config:"server.writeTimeout:60"`
	ReadTimeout  int           `config:"server.readTimeout:60"`
	IdleTimeout  time.Duration `config:"server.idleTimeout:60"`
	Status       bool
}

// NewConfig get Config with reflection automatically
func NewConfig(parser config.Parser) *Config {
	c := Config{}
	config.SetConfig(parser, &c)
	return &c
}

// NewConfig2 get Config manually using parser functions
func NewConfig2(parser config.Parser) *Config {
	t := reflect.TypeOf(Config{})
	return &Config{
		Name:         parser.GetString(config.GetTagValue(t, "Name")),
		LogLevel:     parser.GetString(config.GetTagValue(t, "LogLevel")),
		Port:         parser.GetInt(config.GetTagValue(t, "Port")),
		WriteTimeout: parser.GetInt(config.GetTagValue(t, "WriteTimeout")),
		ReadTimeout:  parser.GetInt(config.GetTagValue(t, "ReadTimeout")),
		IdleTimeout:  time.Duration(parser.GetInt(config.GetTagValue(t, "IdleTimeout"))),
	}
}

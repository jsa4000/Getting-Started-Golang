package server

import (
	"reflect"
	"time"
	"webapp/config"
)

// AppConfig main app configuration
type AppConfig struct {
	Name         string        `config:"app.name:ServerApp"`
	LogLevel     string        `config:"logging.level:info"`
	Port         int           `config:"server.port:8080"`
	WriteTimeout int           `config:"server.writeTimeout:60"`
	ReadTimeout  int           `config:"server.readTimeout:60"`
	IdleTimeout  time.Duration `config:"server.idleTimeout:60"`
}

// NewAppConfig get appconfig
func NewAppConfig(parser config.Parser) *AppConfig {
	t := reflect.TypeOf(AppConfig{})
	return &AppConfig{
		Name:         parser.GetString(config.GetTagValue(t, "Name")),
		LogLevel:     parser.GetString(config.GetTagValue(t, "LogLevel")),
		Port:         parser.GetInt(config.GetTagValue(t, "Port")),
		WriteTimeout: parser.GetInt(config.GetTagValue(t, "WriteTimeout")),
		ReadTimeout:  parser.GetInt(config.GetTagValue(t, "ReadTimeout")),
		IdleTimeout:  time.Duration(parser.GetInt(config.GetTagValue(t, "IdleTimeout"))),
	}
}

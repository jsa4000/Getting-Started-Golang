package viper

import (
	"os"
	"reflect"
	"testing"
	"time"
	"webapp/core/config"
	log "webapp/core/logging"
	"webapp/core/logging/logger"

	"github.com/stretchr/testify/assert"
)

var yamlConfig = []byte(`
app:
  name: WebApp
  version: 1.12

logging:
  enabled: true
  level: debug

server:
  port: 9000
  writeTimeout: 15
  idleTimeout: 200

`)

type Config struct {
	Name         string `config:"app.name:ServerApp"`
	LogLevel     string `config:"logging.level:info"`
	Port         int    `config:"server.port"`
	WriteTimeout int
	ReadTimeout  int           `config:"server.readTimeout:60"`
	IdleTimeout  time.Duration `config:"server.idleTimeout:60"`
	Status       bool
}

var parser *Parser

func init() {
	log.SetGlobal(logger.New())
	// Set the log formatter
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(log.TextFormat)

	parser = NewParserFromBytes(yamlConfig, "yaml")
}

func TestGetStringFromFile(t *testing.T) {
	path := "app.name"
	expectedValue := "WebApp"

	value := parser.GetString(path)

	assert.NotEqual(t, "", value)
	assert.Equal(t, expectedValue, value)
}

func TestGetStringFromFileNotDeafult(t *testing.T) {
	path := "app.name:WebAppDefault"
	expectedValue := "WebApp"

	value := parser.GetString(path)

	assert.NotEqual(t, "", value)
	assert.Equal(t, expectedValue, value)
}

func TestGetFloat64FromFile(t *testing.T) {
	path := "app.version"
	expectedValue := 1.12

	value := parser.GetFloat64(path)

	assert.Equal(t, expectedValue, value)
}

func TestGetFromFile(t *testing.T) {
	path := "logging.level"
	expectedValue := "debug"

	value, err := parser.Get(path)

	assert.Equal(t, expectedValue, value)
	assert.Equal(t, nil, err)
}

func TestGetFromFile2(t *testing.T) {
	path := "logging.enabled"
	expectedValue := true

	value, err := parser.Get(path)

	assert.Equal(t, expectedValue, value)
	assert.Equal(t, nil, err)
}

func TestGetStringFromEnv(t *testing.T) {
	os.Setenv("ENV_APP_NAME", "WebApp")
	path := "env.app.name"
	expectedValue := "WebApp"

	value := parser.GetString(path)

	assert.NotEqual(t, "", value)
	assert.Equal(t, expectedValue, value)

	os.Unsetenv("ENV_APP_NAME")
}

func TestGetStringFromEnvFirst(t *testing.T) {
	os.Setenv("APP_NAME", "WebAppEnv")
	path := "app.name"
	expectedValue := "WebAppEnv"

	value := parser.GetString(path)

	assert.NotEqual(t, "", value)
	assert.Equal(t, expectedValue, value)

	os.Unsetenv("APP_NAME")
}

func TestGetStringFromDefault(t *testing.T) {
	path := "app.default:WebAppDefault"
	expectedValue := "WebAppDefault"

	value := parser.GetString(path)

	assert.NotEqual(t, "", value)
	assert.Equal(t, expectedValue, value)
}

func TestGetIntFromDefault(t *testing.T) {
	path := "app.integer:1"
	expectedValue := 1

	value := parser.GetInt(path)

	assert.Equal(t, expectedValue, value)
}

func TestGetFloat64FromDefault(t *testing.T) {
	path := "app.float:1.12"
	expectedValue := 1.12

	value := parser.GetFloat64(path)

	assert.Equal(t, expectedValue, value)
}

func TestGetBoolFromDefault(t *testing.T) {
	path := "app.bool:true"
	expectedValue := true

	value := parser.GetBool(path)

	assert.Equal(t, expectedValue, value)
}

func TestGetStringFromFileNotDeafult2(t *testing.T) {
	path := "logging.level:info"
	expectedValue := "debug"

	value := parser.GetString(path)

	assert.NotEqual(t, "", value)
	assert.Equal(t, expectedValue, value)
}

func TestGetError(t *testing.T) {
	path := "app.fail"
	_, err := parser.Get(path)

	assert.NotEqual(t, nil, err)
}

// NewConfig get Config with reflection automatically
func TestParseFieldsByReflection(t *testing.T) {
	c := Config{}
	parser.ReadFields(&c)

	assert.Equal(t, c.Name, "WebApp")
	assert.Equal(t, c.LogLevel, "debug")
	assert.Equal(t, c.ReadTimeout, 60)
	assert.Equal(t, c.IdleTimeout, time.Duration(200))
}

// NewConfig2 get Config manually using parser functions
func TestParseFieldsManually(t *testing.T) {
	rt := reflect.TypeOf(Config{})
	c := &Config{
		Name:         parser.GetString(config.GetTagValue(rt, "Name")),
		LogLevel:     parser.GetString(config.GetTagValue(rt, "LogLevel")),
		Port:         parser.GetInt(config.GetTagValue(rt, "Port")),
		WriteTimeout: parser.GetInt(config.GetTagValue(rt, "WriteTimeout")),
		ReadTimeout:  parser.GetInt(config.GetTagValue(rt, "ReadTimeout")),
		IdleTimeout:  time.Duration(parser.GetInt(config.GetTagValue(rt, "IdleTimeout"))),
	}

	assert.Equal(t, c.Name, "WebApp")
	assert.Equal(t, c.LogLevel, "debug")
	assert.Equal(t, c.ReadTimeout, 60)
	assert.Equal(t, c.IdleTimeout, time.Duration(200))
}

package viper

import (
	"os"
	"reflect"
	"testing"
	"time"
	"webapp/core/config"

	"github.com/stretchr/testify/assert"
)

var yamlConfig = []byte(`
app:
  name: WebApp
  version: 1.12
    
logging:
  enabled: true
  level: debug
`)

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

var parser *Parser

func init() {
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
func NewConfig(parser config.Parser) *Config {
	c := Config{}
	//config.SetConfig(parser, &c)
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

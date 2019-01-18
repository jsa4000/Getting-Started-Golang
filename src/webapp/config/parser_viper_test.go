package config

import (
	"os"
	"testing"

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

var parser Parser

func init() {
	parser = NewViperParserFromBytes(yamlConfig, "yaml")
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

// To test when override previous properties with Viper.Set()
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

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

	assert.NotEqual(t, value, "")
	assert.Equal(t, value, expectedValue)
}

func TestGetFloat64FromFile(t *testing.T) {
	path := "app.version"
	expectedValue := 1.12

	value := parser.GetFloat64(path)

	assert.Equal(t, value, expectedValue)
}

func TestGetFromFile(t *testing.T) {
	path := "logging.level"
	expectedValue := "debug"

	value, err := parser.Get(path)

	assert.Equal(t, value, expectedValue)
	assert.Equal(t, err, nil)
}

func TestGetFromFile2(t *testing.T) {
	path := "logging.enabled"
	expectedValue := true

	value, err := parser.Get(path)

	assert.Equal(t, value, expectedValue)
	assert.Equal(t, err, nil)
}

func TestGetStringFromEnv(t *testing.T) {
	os.Setenv("ENV_APP_NAME", "WebApp")
	path := "env.app.name"
	expectedValue := "WebApp"

	value := parser.GetString(path)

	assert.NotEqual(t, value, "")
	assert.Equal(t, value, expectedValue)

	os.Unsetenv("ENV_APP_NAME")
}

func TestGetStringFromEnvFirst(t *testing.T) {
	os.Setenv("APP_NAME", "WebAppEnv")
	path := "app.name"
	expectedValue := "WebAppEnv"

	value := parser.GetString(path)

	assert.NotEqual(t, value, "")
	assert.Equal(t, value, expectedValue)

	os.Unsetenv("APP_NAME")
}

func TestGetStringError(t *testing.T) {
	path := "app.fail"
	value := parser.GetString(path)

	assert.Equal(t, value, "")
}

func TestGetError(t *testing.T) {
	path := "app.fail"
	_, err := parser.Get(path)

	assert.NotEqual(t, err, nil)
}

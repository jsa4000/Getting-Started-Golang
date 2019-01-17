package config

import (
	"testing"
)

// GetPathFromConfigfileShouldReturnOk
func TestReadPath(t *testing.T) {
	parser := NewViperParser("config_test", ".")
	path := "app.name"
	value, err := parser.GetString(path)

	expectedValue := "WebApp"

	if err != nil || value == "" {
		t.Errorf("Error getting %s", path)
	}

	if value != expectedValue {
		t.Errorf("test error %s != %s ", expectedValue, value)
	}
}

// GetPathFromConfigfileShouldReturnError
func TestReadPathError(t *testing.T) {
	parser := NewViperParser("config_test", ".")
	path := "app.fail"
	value, err := parser.GetString(path)

	if err == nil || value != "" {
		t.Errorf("Path must not exist %s", path)
	}
}

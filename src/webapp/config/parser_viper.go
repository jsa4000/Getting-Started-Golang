package config

import (
	"fmt"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ViperParser structure
type ViperParser struct {
}

// NewViperParser creates the default parser implementation
func NewViperParser(filename string, path string) Parser {
	parser := ViperParser{}

	viper.SetConfigName(strings.TrimSuffix(filename, filepath.Ext(filename)))
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return &parser
}

// GetString from a path
func (p *ViperParser) GetString(path string) (string, error) {
	value := viper.GetString(path)
	if value == "" {
		return "", fmt.Errorf("Path not found: %s", path)
	}
	return value, nil
}

// Get from a path
func (p *ViperParser) Get(path string) (interface{}, error) {
	value := viper.Get(path)
	if value == "" {
		return "", fmt.Errorf("Path not found: %s", path)
	}
	return value, nil
}

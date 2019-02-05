package viper

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"webapp/core/config"

	"github.com/spf13/viper"
)

// Parser structure
type Parser struct {
}

// New creates the default parser implementation
func New() *Parser {
	viper.New()
	return &Parser{}
}

// LoadFromFile creates the default parser implementation
func (p *Parser) LoadFromFile(filename string, path string) error {
	//viper.SetConfigType("yaml") // Inferred
	viper.SetConfigName(strings.TrimSuffix(filename, filepath.Ext(filename)))
	viper.AddConfigPath(path)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// LoadFromBytes creates the default parser implementation
func (p *Parser) LoadFromBytes(buffer []byte, filetype string) error {
	viper.SetConfigType(filetype)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadConfig(bytes.NewBuffer(buffer)); err != nil {
		return err
	}
	return nil
}

// ReadFields read fields tags from struct and return config values
func (p *Parser) ReadFields(data interface{}) {
	config.ReadData(p, data)
}

// GetString from a path
func (p *Parser) GetString(path string) string {
	return viper.GetString(config.ProcessPath(p, path))
}

// GetFloat64 from a path
func (p *Parser) GetFloat64(path string) float64 {
	return viper.GetFloat64(config.ProcessPath(p, path))
}

// GetBool from a path
func (p *Parser) GetBool(path string) bool {
	return viper.GetBool(config.ProcessPath(p, path))
}

// GetInt from a path
func (p *Parser) GetInt(path string) int {
	return viper.GetInt(config.ProcessPath(p, path))
}

// Get from a path
func (p *Parser) Get(path string) (interface{}, error) {
	value := viper.Get(config.ProcessPath(p, path))
	if value == nil {
		return nil, fmt.Errorf("Path not found: %s", path)
	}
	return value, nil
}

// SetDefault value when a variable is not configured
func (p *Parser) SetDefault(key string, value interface{}) {
	if viper.Get(key) == nil {
		viper.Set(key, value)
	}
}

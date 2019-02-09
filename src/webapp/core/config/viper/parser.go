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
	Viper *viper.Viper
}

// New creates the default parser implementation
func New() *Parser {
	return &Parser{
		Viper: viper.New(),
	}
}

// LoadFromFile creates the default parser implementation
func (p *Parser) LoadFromFile(filename string, path string) error {
	//p.Viper.SetConfigType("yaml") // Inferred
	p.Viper.SetConfigName(strings.TrimSuffix(filename, filepath.Ext(filename)))
	p.Viper.AddConfigPath(path)
	p.defaults()
	if err := p.Viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// LoadFromBytes creates the default parser implementation
func (p *Parser) LoadFromBytes(buffer []byte, filetype string) error {
	p.Viper.SetConfigType(filetype)
	p.defaults()
	if err := p.Viper.ReadConfig(bytes.NewBuffer(buffer)); err != nil {
		return err
	}
	return nil
}

func (p *Parser) defaults() {
	p.Viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	p.Viper.AutomaticEnv()
}

// UnmarshalKey serialize current val from configuration
func (p *Parser) UnmarshalKey(key string, val interface{}) error {
	return p.Viper.UnmarshalKey(key, val)
}

// ReadFields read fields tags from struct and return config values
func (p *Parser) ReadFields(data interface{}) {
	config.ReadData(p, data)
}

// GetString from a path
func (p *Parser) GetString(path string) string {
	return p.Viper.GetString(config.ProcessPath(p, path))
}

// GetFloat64 from a path
func (p *Parser) GetFloat64(path string) float64 {
	return p.Viper.GetFloat64(config.ProcessPath(p, path))
}

// GetBool from a path
func (p *Parser) GetBool(path string) bool {
	return p.Viper.GetBool(config.ProcessPath(p, path))
}

// GetInt from a path
func (p *Parser) GetInt(path string) int {
	return p.Viper.GetInt(config.ProcessPath(p, path))
}

// Get from a path
func (p *Parser) Get(path string) (interface{}, error) {
	value := p.Viper.Get(config.ProcessPath(p, path))
	if value == nil {
		return nil, fmt.Errorf("Path not found: %s", path)
	}
	return value, nil
}

// SetDefault value when a variable is not configured
func (p *Parser) SetDefault(key string, value interface{}) {
	if p.Viper.Get(key) == nil {
		p.Viper.Set(key, value)
	}
}

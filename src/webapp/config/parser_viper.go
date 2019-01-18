package config

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ViperParser structure
type ViperParser struct {
}

// NewViperParserFromFile creates the default parser implementation
func NewViperParserFromFile(filename string, path string) Parser {
	parser := ViperParser{}

	//viper.SetConfigType("yaml") // Inferred
	viper.SetConfigName(strings.TrimSuffix(filename, filepath.Ext(filename)))
	viper.AddConfigPath(path)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	return &parser
}

// NewViperParserFromBytes creates the default parser implementation
func NewViperParserFromBytes(buffer []byte, filetype string) Parser {
	parser := ViperParser{}

	viper.New()
	viper.SetConfigType(filetype)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadConfig(bytes.NewBuffer(buffer)); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	return &parser
}

//preProcessDefaultsInPath
func (p *ViperParser) preProcessDefaultsInPath(path string) string {
	items := strings.SplitN(path, ":", 2)
	if len(items) == 1 {
		return items[0]
	}
	p.SetDefault(items[0], items[1])
	return items[0]
}

// GetString from a path
func (p *ViperParser) GetString(path string) string {
	return viper.GetString(p.preProcessDefaultsInPath(path))
}

// GetFloat64 from a path
func (p *ViperParser) GetFloat64(path string) float64 {
	return viper.GetFloat64(p.preProcessDefaultsInPath(path))
}

// GetBool from a path
func (p *ViperParser) GetBool(path string) bool {
	return viper.GetBool(p.preProcessDefaultsInPath(path))
}

// GetInt from a path
func (p *ViperParser) GetInt(path string) int {
	return viper.GetInt(p.preProcessDefaultsInPath(path))
}

// Get from a path
func (p *ViperParser) Get(path string) (interface{}, error) {
	value := viper.Get(p.preProcessDefaultsInPath(path))
	if value == nil {
		return nil, fmt.Errorf("Path not found: %s", path)
	}
	return value, nil
}

// SetDefault value when a variable is not configured
func (p *ViperParser) SetDefault(key string, value interface{}) {
	if viper.Get(key) == nil {
		viper.Set(key, value)
	}
}

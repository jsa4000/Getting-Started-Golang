package config

import "reflect"

// Tag used for inference path and default value
// i.e  Name   string  `config:"app.name:ServerApp"`
const Tag = "config"

// Parser interface to read from config files or environment
type Parser interface {
	SetDefault(key string, value interface{})
	Get(path string) (interface{}, error)
	GetInt(path string) int
	GetBool(path string) bool
	GetFloat64(path string) float64
	GetString(path string) string
}

// GetTagValue gets a tag from a field and from a Type
func GetTagValue(t reflect.Type, field string) string {
	f, _ := t.FieldByName(field)
	value, _ := f.Tag.Lookup(Tag)
	return value
}

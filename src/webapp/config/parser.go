package config

// Parser interface to read from config files or environment
type Parser interface {
	SetDefault(key string, value interface{})
	Get(path string) (interface{}, error)
	GetInt(path string) int
	GetBool(path string) bool
	GetFloat64(path string) float64
	GetString(path string) string
}

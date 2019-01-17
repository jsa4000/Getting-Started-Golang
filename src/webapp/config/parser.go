package config

// Parser interface to read from config files or environment
type Parser interface {
	Get(path string) (interface{}, error)
	GetString(path string) (string, error)
}

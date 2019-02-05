package config

const (
	// Tag used for inference path and default value
	// i.e  Name   string  `config:"app.name:ServerApp"`
	Tag = "config"
	// PathSeparator to split the tag into `config:"path:default"`
	PathSeparator = ":"
)

// Parser interface to read from config files or environment
// path must be in the form 'configPath:defaultValue'. i.e "app.name:MyServerApp"
type Parser interface {
	LoadFromFile(filename string, path string) error
	LoadFromBytes(buffer []byte, filetype string) error

	ReadFields(data interface{})
	Get(path string) (interface{}, error)
	GetInt(path string) int
	GetBool(path string) bool
	GetFloat64(path string) float64
	GetString(path string) string
	SetDefault(key string, value interface{})
}

// Config Global logger
var parser Parser

// SetGlobal sets the Global Logger (singletone)
func SetGlobal(p Parser) {
	parser = p
}

// ReadFields read fields tags from struct and return config values
func ReadFields(data interface{}) {
	parser.ReadFields(data)
}

// GetString from a path
func GetString(path string) string {
	return parser.GetString(path)
}

// GetFloat64 from a path
func GetFloat64(path string) float64 {
	return parser.GetFloat64(path)
}

// GetBool from a path
func GetBool(path string) bool {
	return parser.GetBool(path)
}

// GetInt from a path
func GetInt(path string) int {
	return parser.GetInt(path)
}

// Get from a path
func Get(path string) (interface{}, error) {
	return parser.Get(path)
}

// SetDefault value when a variable is not configured
func SetDefault(key string, value interface{}) {
	parser.SetDefault(key, value)
}

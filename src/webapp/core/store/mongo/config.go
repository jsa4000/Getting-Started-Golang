package mongo

import "webapp/core/config"

// Config main app configuration
type Config struct {
	URL string `config:"repository.mongodb.url:mongodb://root:root@localhost:27017/admin"`
}

// LoadConfig Load config from file
func LoadConfig() *Config {
	c := Config{}
	config.ReadFields(&c)
	return &c
}

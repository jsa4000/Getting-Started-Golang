package redis

import (
	"time"
	"webapp/core/config"
)

// Config main app configuration
type Config struct {
	URL            string        `config:"cache.redis.url:redis://:root@dockerhost:6379/0"`
	Addrs          string        `config:"cache.redis.addrs"`
	Password       string        `config:"cache.redis.password"`
	Database       int           `config:"cache.redis.database:0"`
	MaxRetries     int           `config:"cache.redis.maxRetries:-1"`
	DialTimeout    time.Duration `config:"cache.redis.dialTimeout:-1"`
	ReadTimeout    time.Duration `config:"cache.redis.readTimeout:-1"`
	WriteTimeout   time.Duration `config:"cache.redis.writeTimeout:-1"`
	PoolSize       int           `config:"cache.redis.poolSize:-1"`
	MaxConnAge     time.Duration `config:"cache.redis.maxConnAge:-1"`
	ReadOnly       bool          `config:"cache.redis.cluster.readOnly"`
	RouteByLatency bool          `config:"cache.redis.cluster.routeByLatency"`
	RouteRandomly  bool          `config:"cache.redis.cluster.routeRandomly"`
	MasterName     string        `config:"cache.redis.cluster.masterName:"`
}

// LoadConfig Load config from file
func LoadConfig() *Config {
	c := Config{}
	config.ReadFields(&c)
	return &c
}

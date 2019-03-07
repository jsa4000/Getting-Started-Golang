package redis

import (
	"context"
	"fmt"
	"time"
	log "webapp/core/logging"

	"github.com/go-redis/redis"
)

// OnConnect Handler that will be called every time is connects
type OnConnect func()

// Wrapper structure for Redis Client
type Wrapper struct {
	Client    *redis.Client
	Config    *Config
	OnConnect map[string]OnConnect
}

var wrapper *Wrapper

// SetGlobal set global component component
func SetGlobal(w *Wrapper) {
	wrapper = w
}

// New returns new Redis client
func New() *Wrapper {
	return &Wrapper{
		OnConnect: map[string]OnConnect{},
	}
}

// Client Returns Redis Client
func Client() *redis.Client {
	return wrapper.Client
}

// Connect to Redis database
func Connect(ctx context.Context, config *Config) error {
	return wrapper.Connect(ctx, config)
}

// Disconnect to Redis database
func Disconnect(ctx context.Context) error {
	return wrapper.Disconnect(ctx)
}

// Set value into the cache
func Set(key string, value interface{}, exp time.Duration) error {
	return wrapper.Set(key, value, exp)
}

// Get the value for the given key if exists
func Get(key string) (string, error) {
	return wrapper.Get(key)
}

// Float64 the value for the given key if exists
func Float64(key string) (float64, error) {
	return wrapper.Float64(key)
}

// Int64 the value for the given key if exists
func Int64(key string) (int64, error) {
	return wrapper.Int64(key)
}

// Connected check the connection to Redis database
func (w *Wrapper) Connected() bool {
	_, err := w.Client.Ping().Result()
	if err != nil {
		return false
	}
	return true
}

// check the connection to Mongodb database
func (w *Wrapper) check(ctx context.Context) {
	for {
		_, err := w.Client.Ping().Result()
		if err != nil {
			log.Error(fmt.Sprintf("Error trying to connect to redis. '%s'", err))
			time.Sleep(5 * time.Second)
			continue
		}
		log.Info(fmt.Sprintf("Connected to redis '%s'", w.Config.URL))
		break
	}
}

// Connect to Redis database
func (w *Wrapper) Connect(ctx context.Context, config *Config) error {
	w.Config = config
	w.Client = redis.NewClient(&redis.Options{
		Addr:       config.URL,
		Password:   config.Password,
		DB:         config.Database,
		MaxRetries: config.MaxRetries,
	})
	go w.check(ctx)
	return nil
}

// Set value into the cache
func (w *Wrapper) Set(key string, value interface{}, exp time.Duration) error {
	return w.Client.Set(key, value, exp).Err()
}

// Get the value for the given key if exists
func (w *Wrapper) Get(key string) (string, error) {
	return w.Client.Get(key).Result()
}

// Float64 the value for the given key if exists
func (w *Wrapper) Float64(key string) (float64, error) {
	return w.Client.Get(key).Float64()
}

// Int64 the value for the given key if exists
func (w *Wrapper) Int64(key string) (int64, error) {
	return w.Client.Get(key).Int64()
}

// GetString the value for the given key if exists
func (w *Wrapper) String(key string) (string, error) {
	return w.Client.Get(key).Result()
}

// Disconnect to Redis database
func (w *Wrapper) Disconnect(ctx context.Context) error {
	err := w.Client.Close()
	if err != nil {
		log.Error(fmt.Sprintf("Error disconnecting from Redis. '%s'", err))
		return err
	}
	log.Info("Connection to Redis closed.")
	return nil
}

package cache

import "time"

// Cache interface
type Cache interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
	Float64(key string) (float64, error)
	Int64(key string) (int64, error)
}

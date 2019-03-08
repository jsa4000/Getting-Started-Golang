package cache

import "time"

var cache Cache

// SetGlobal set global component component
func SetGlobal(c Cache) {
	cache = c
}

// Set value into the cache
func Set(key string, value interface{}, exp time.Duration) error {
	return cache.Set(key, value, exp)
}

// Get returns the value for the given key if exists
func Get(key string) (string, error) {
	return cache.Get(key)
}

// Delete the value for the given key if exists
func Delete(key string) error {
	return cache.Delete(key)
}

// Bytes returns the value for the given key if exists
func Bytes(key string) ([]byte, error) {
	return cache.Bytes(key)
}

// Float64 returns the value for the given key if exists
func Float64(key string) (float64, error) {
	return cache.Float64(key)
}

// Int64 returns the value for the given key if exists
func Int64(key string) (int64, error) {
	return cache.Int64(key)
}

package roles

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"webapp/core/config"
	log "webapp/core/logging"
	"webapp/core/store/cache"
)

var cacheTTL = 3600 * time.Second

const (
	cacheID       = "Roles"
	cacheFindAll  = "FindAll"
	cacheFindByID = "FindByID"
)

// CacheConfig for mongoDB Repository
type CacheConfig struct {
}

// CacheRepository to implement the Roles Repository
type CacheRepository struct {
	repository Repository
}

// NewCacheRepository Create a Mock repository
func NewCacheRepository(repository Repository) Repository {
	c := CacheConfig{}
	config.ReadFields(&c)
	result := &CacheRepository{
		repository: repository,
	}
	return result
}

func key(function string, params ...string) string {
	if len(params) == 0 {
		return fmt.Sprintf("%s;%s", cacheID, function)
	}
	return fmt.Sprintf("%s;%s;%s", cacheID, function, strings.Join(params, ":"))
}

// FindAll fetches all the values form the database
func (c *CacheRepository) FindAll(ctx context.Context) ([]*Role, error) {
	cachekey := key(cacheFindAll)
	value, err := cache.Bytes(cachekey)
	if err == nil {
		var result []*Role
		if err := json.Unmarshal(value, &result); err == nil {
			log.Debugf("Used cache stored in '%s'", cachekey)
			return result, nil
		}
	}
	result, err := c.repository.FindAll(ctx)
	if err == nil {
		if bytes, err2 := json.Marshal(result); err2 == nil {
			cache.Set(cachekey, bytes, cacheTTL)
		}
	}
	return result, err
}

// FindByID Role by Id
func (c *CacheRepository) FindByID(ctx context.Context, id string) (*Role, error) {
	cachekey := key(cacheFindByID, id)
	value, err := cache.Bytes(cachekey)
	if err == nil {
		var result Role
		if err := json.Unmarshal(value, &result); err == nil {
			log.Debugf("Used cache stored in '%s'", cachekey)
			return &result, nil
		}
	}
	result, err := c.repository.FindByID(ctx, id)
	if err == nil {
		if bytes, err2 := json.Marshal(result); err2 == nil {
			cache.Set(cachekey, bytes, cacheTTL)
		}
	}
	return result, err
}

// Create Add Role into the datbase
func (c *CacheRepository) Create(ctx context.Context, Role Role) (*Role, error) {
	result, err := c.repository.Create(ctx, Role)
	if err == nil {
		cache.Delete(key(cacheFindAll))
	}
	return result, err
}

// DeleteByID Role from the database
func (c *CacheRepository) DeleteByID(ctx context.Context, id string) (bool, error) {
	deleted, err := c.repository.DeleteByID(ctx, id)
	if deleted {
		cache.Delete(key(cacheFindAll))
		cache.Delete(key(cacheFindByID, id))
	}
	return deleted, err
}

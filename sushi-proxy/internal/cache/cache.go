package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// DefaultExpiration is a helper value that uses the cache defaults.
const DEFAULT_EXPIRATION int64 = 0

type CacheImpl interface {
	Get(string) (interface{}, bool)
	Set(string, interface{}, int64)
	Delete(string)
	Count() int
	Flush()
}

type BasicCache struct {
	// Default expiration interval, in seconds.
	defaultExpiration int64

	// Clean up interval, in seconds.
	cleanupInterval int64

	// The underlying cache driver.
	cache *cache.Cache
}

// New creates a new cache instance.
func New(defaultExpiration int64, cleanupInterval int64) *BasicCache {
	defaultExpirationSecs := time.Duration(defaultExpiration) * time.Second
	cleanupIntervalSecs := time.Duration(cleanupInterval) * time.Second

	return &BasicCache{
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
		cache:             cache.New(defaultExpirationSecs, cleanupIntervalSecs),
	}
}

// Get retrieves a cache item by key.
func (bc *BasicCache) Get(key string) (interface{}, bool) {
	return bc.cache.Get(key)
}

// Set writes a cache item with a timeout in seconds. If timeout is zero,
// the default expiration for the repository instance will be used.
func (bc *BasicCache) Set(key string, value interface{}, timeout int64) {
	if timeout <= 0 {
		timeout = bc.defaultExpiration
	}
	bc.cache.Set(key, value, time.Duration(timeout)*time.Second)
}

// Delete cache item by key.
func (bc *BasicCache) Delete(key string) {
	bc.cache.Delete(key)
}

// Count returns number of items in the cache.
func (bc *BasicCache) Count() int {
	return bc.cache.ItemCount()
}

// Flush flushes all the items from the cache.
func (bc *BasicCache) Flush() {
	bc.cache.Flush()
}

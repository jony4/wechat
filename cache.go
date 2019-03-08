package wechat

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache interface for store access token.
type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, expiration time.Duration)
	Delete(key string)
}

// const of cache
const (
	DefaultInterval   time.Duration = 1 * time.Minute // min unit set to ms
	DefaultExpiration time.Duration = 0
	NoExpiration      time.Duration = -1
)

// MemCache MemCache
type MemCache struct {
	cache      *cache.Cache
	expiration time.Duration
	interval   time.Duration
}

// MemCacheOptFunc MemCacheOptFunc
type MemCacheOptFunc func(*MemCache) error

// NewMemCache NewMemCache
func NewMemCache(options ...MemCacheOptFunc) (*MemCache, error) {
	mc := &MemCache{
		expiration: DefaultExpiration,
		interval:   DefaultInterval,
	}
	for _, option := range options {
		if err := option(mc); err != nil {
			return nil, err
		}
	}
	mc.cache = cache.New(mc.expiration, mc.interval)
	return mc, nil
}

// SetDefaultExpiration SetDefaultExpiration
func SetDefaultExpiration(exp time.Duration) MemCacheOptFunc {
	return func(cache *MemCache) error {
		cache.expiration = exp
		return nil
	}
}

// SetDefaultInterval SetDefaultInterval
func SetDefaultInterval(inter time.Duration) MemCacheOptFunc {
	return func(cache *MemCache) error {
		cache.interval = inter
		return nil
	}
}

// Set Set
func (mc *MemCache) Set(key string, value interface{}, expiration time.Duration) {
	mc.cache.Set(key, value, expiration)
}

// Get Get
func (mc *MemCache) Get(key string) (interface{}, bool) {
	return mc.cache.Get(key)
}

// Delete Delete
func (mc *MemCache) Delete(key string) {
	mc.cache.Delete(key)
}

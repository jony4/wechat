package wechat

import (
	"context"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache interface for store access token.
type Cache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
}

// const of cache
const (
	DefaultInterval   time.Duration = 1 * time.Minute // min unit set to ms
	DefaultExpiration time.Duration = 0
	NoExpiration      time.Duration = -1
)

// var of cache
var (
	ErrCacheKeyNotExist = errors.New("key not exist")
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
func (mc *MemCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	c := make(chan struct{}, 1)
	go func() {
		mc.cache.Set(key, value, expiration)
		c <- struct{}{}
	}()
	select {
	case <-c:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Get Get
func (mc *MemCache) Get(ctx context.Context, key string) (value interface{}, err error) {
	cs := make(chan struct{}, 1)
	ci := make(chan interface{}, 1)
	go func() {
		vv, exist := mc.cache.Get(key)
		if !exist {
			cs <- struct{}{}
			return
		}
		ci <- vv
	}()
	select {
	case v := <-ci:
		value = v
	case <-cs:
		return nil, ErrCacheKeyNotExist
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return value, nil
}

// Delete Delete
func (mc *MemCache) Delete(ctx context.Context, key string) error {
	c := make(chan struct{}, 1)
	go func() {
		mc.cache.Delete(key)
		c <- struct{}{}
	}()
	select {
	case <-c:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

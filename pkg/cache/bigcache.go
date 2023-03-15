package cache

import (
	"time"

	"github.com/allegro/bigcache"
)

type BigCache struct {
	cache *bigcache.BigCache
}

func NewBigCache() *BigCache {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Second * 3600))
	return &BigCache{
		cache: cache,
	}
}

func (c *BigCache) Get(key string) ([]byte, error) {
	return c.cache.Get(key)
}

func (c *BigCache) Set(key string, value []byte) error {
	return c.cache.Set(key, value)
}

func (c *BigCache) Delete(key string) error {
	return c.cache.Delete(key)
}

func (c *BigCache) Clear() error {
	return c.cache.Reset()
}

func (c *BigCache) GetType() string {
	return "BigCache"
}

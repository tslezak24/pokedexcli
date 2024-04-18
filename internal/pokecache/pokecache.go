package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{entry: make(map[string]cacheEntry), mux: &sync.Mutex{}}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) AddToCache(key string, value []byte) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	cache.entry[key] = cacheEntry{createdAt: time.Now().UTC(), val: value}
}
func (cache *Cache) GetFromCache(key string) ([]byte, bool) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	val, ok := cache.entry[key]
	return val.val, ok
}
func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.entry {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entry, k)
		}
	}
}

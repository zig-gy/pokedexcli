package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	mu 		sync.Mutex
	entries	map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	newC :=  &Cache{
		entries: map[string]cacheEntry{},
	}

	go newC.reapLoop(interval)
	return newC
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.entries[key] = entry
}

func (c *Cache) Get(key string) (entry []byte, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ent, ok := c.entries[key]
	if !ok {
		return
	}

	entry = ent.val
	return
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()

		for k, v := range c.entries {
			timeDiff := time.Since(v.createdAt)
			if  timeDiff > interval{
				delete(c.entries, k)
			}
		}

		c.mu.Unlock()
	}
}
package internal

import (
	"sync"
	"time"
  // "fmt"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	cachedData, exists := c.entry[key]
	return cachedData.val, exists
}

func NewCache(interval time.Duration) Cache {
	cacheElement := Cache{
		entry: map[string]cacheEntry{},
		mu:    &sync.Mutex{},
	}
	go cacheElement.reaploop(interval)
	return cacheElement
}

func (c Cache) reaploop(interval time.Duration) {
	tk := time.NewTicker(interval)
	for range tk.C {
		c.mu.Lock()
    // fmt.Println("Checking cache")
		for key, cacheEntryValue := range c.entry {
			if time.Since(cacheEntryValue.createdAt) > interval {
				delete(c.entry, key)
        // fmt.Println(key)
			}
		}
		c.mu.Unlock()
	}
}

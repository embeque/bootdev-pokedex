package pokecache

import (
	"sync"
	"time"
)

var mu sync.Mutex

type Cache struct {
	data map[string]CachedEntry
	// mu *sync.Mutex
}
type CachedEntry struct {
	createdAt time.Time
	val       []byte
}

func reapLoop(interval time.Duration, cache *Cache) {
	chann := time.Tick(interval)

	for range chann {
		for key := range cache.data {
			if cache.data[key].createdAt.Before(time.Now().Add(-interval)) {
				delete(cache.data, key)
			}
		}
	}

}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data: make(map[string]CachedEntry),
		// mu: &sync.Mutex{}
	}

	go reapLoop(interval, &cache)

	return cache
}

func (c Cache) Add(key string, value []byte) error {
	entry := CachedEntry{
		createdAt: time.Now(),
		val:       value,
	}
	mu.Lock()
	defer mu.Unlock()

	c.data[key] = entry
	return nil

}

func (c Cache) Get(key string) ([]byte, bool) {
	mu.Lock()
	defer mu.Unlock()

	entry, ok := c.data[key]
	return entry.val, ok
}

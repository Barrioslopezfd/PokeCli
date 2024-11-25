package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
    createAt    time.Time
    val         []byte
}

type Cache struct{
    cache       map[string]cacheEntry
    mu          sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
    ch := &Cache{
        cache: make(map[string]cacheEntry),
    }

    ch.reapLoop(interval)

    return ch
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    item := c.cache[key]
    item.val = val
    c.cache[key] = item
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    val, ok := c.cache[key]
    if ok {
        return val.val, true
    }
    return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    go func() {
        for range ticker.C {
            c.mu.Lock()
            for k:= range c.cache {
                if time.Since(c.cache[k].createAt) > interval {
                    delete(c.cache, k)
                }
            }
            c.mu.Unlock()
        }
    }()
}


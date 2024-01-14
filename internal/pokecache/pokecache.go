package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
  createdAt time.Time
  val []byte
}

type Cache struct {
  c map[string]cacheEntry
  mu *sync.Mutex
  interval time.Duration
}

func NewCache(interval time.Duration) Cache {
  c :=  Cache{
    c: make(map[string]cacheEntry),
    mu: &sync.Mutex{},
  }
  c.reapLoop(interval)
  return c
}

func (c *Cache) Add(key string, val []byte) {
  c.mu.Lock()
  defer c.mu.Unlock()
  c.c[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
  c.mu.Lock()
  defer c.mu.Unlock()
  cacheEntry, ok := c.c[key]
  return cacheEntry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
  ticker := time.NewTicker(interval)
  go func() {
    for t := range ticker.C {
      // remove all entries older than interval
      for k, v := range c.c {
        if t.Sub(v.createdAt) > interval {
          delete(c.c, k)
        }
      }
    }
  }()
}

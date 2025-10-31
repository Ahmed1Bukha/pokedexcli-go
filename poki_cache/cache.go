package poki_cache

import (
	"sync"
	"time"
)

type Cache struct{
	entries map[string]CacheEntry
	mu sync.Mutex
	interval time.Duration
}

type CacheEntry struct{
	createAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache{
	c := &Cache{
		entries: make(map[string]CacheEntry),
		interval: interval,
	}
	go func(){
		for range interval/2{
			c.reapLoob()
		}
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		createAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte,bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	entry:=  c.entries[key]
	if entry.val ==nil{
		return nil,false
	}
	return entry.val,true
}

func (c *Cache) reapLoob(){
	c.mu.Lock()
	defer c.mu.Unlock()
	for key,entry := range c.entries{
		if time.Since(entry.createAt) > c.interval{
			delete(c.entries,key)
		}
	}
}
package simpleCache

import (
	"fmt"
	"sync"
	"time"
)

type Cashier interface {
	Set(key string, value interface{}, dur time.Duration)
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type Cache struct {
	storage map[string]CacheItem
	sync.RWMutex
}

type CacheItem struct {
	value interface{}
	ts    time.Time
	dur   time.Duration
}

func NewCache() *Cache {
	return &Cache{storage: make(map[string]CacheItem)}
}

func (c *Cache) Set(key string, value interface{}, dur time.Duration) {
	c.Lock()
	defer c.Unlock()
	c.storage[key] = CacheItem{
		value: value,
		ts:    time.Now(),
		dur:   dur,
	}
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.Lock()
	defer c.Unlock()
	if item, ok := c.storage[key]; ok {
		if time.Since(item.ts) < item.dur {
			return item.value, nil
		} else {
			delete(c.storage, key)
			return nil, fmt.Errorf("no data found")
		}
	}
	return nil, fmt.Errorf("no data found")
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.storage[key]; ok {
		delete(c.storage, key)
	}
	return fmt.Errorf("no data found")
}

package cache

import (
	"sync"
)

type Cache struct {
	store map[string]any
}

var (
	instance *Cache
	once     sync.Once
)

func New() *Cache {
	once.Do(func() {
		instance = &Cache{store: make(map[string]any)}
	})

	return instance
}

func (c *Cache) Set(key string, value any) {
	c.store[key] = value
}

func (c *Cache) Get(key string, value any) any {
	return c.store[key]
}

func (c *Cache) Delete(key string) {
	_, ok := c.store[key]

	if !ok {
		return
	}

	delete(c.store, key)
}

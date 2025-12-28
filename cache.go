package cache

import (
	"sync"
)

type Cache struct {
	store map[string]interface{}
}

var (
	instance *Cache
	once     sync.Once
)

func New() *Cache {
	once.Do(func() {
		instance = &Cache{store: make(map[string]interface{})}
	})

	return instance
}

func (c *Cache) Set(key string, value interface{}) {
	c.store[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.store[key]
}

func (c *Cache) Delete(key string) {
	_, ok := c.store[key]

	if !ok {
		return
	}

	delete(c.store, key)
}

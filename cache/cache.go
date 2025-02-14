package cache

import (
	"sync"

	"github.com/cycloidio/terracognita/errcode"
	"github.com/cycloidio/terracognita/provider"
)

// Cache implements a simple cache of provider.Resource
// it's not concurrently safe
type Cache interface {
	// Set set's the rs to the key
	// if an already existing key
	// was there, it'll return an error
	Set(key string, rs []provider.Resource) error

	// Get get's the values of the key
	// if the key is not found an error
	// is returned
	Get(key string) ([]provider.Resource, error)
}

type cache struct {
	lock sync.Mutex
	data map[string][]provider.Resource
}

// New returns a new Cache implementaion
func New() Cache {
	return &cache{
		lock: sync.Mutex{},
		data: make(map[string][]provider.Resource),
	}
}

func (c *cache) Set(key string, rs []provider.Resource) error {
	val, ok := c.data[key]
	if ok {
		rs = append(val, rs...)
	}
	c.lock.Lock()
	c.data[key] = rs
	c.lock.Unlock()
	return nil
}

func (c *cache) Get(key string) ([]provider.Resource, error) {
	c.lock.Lock()
	rs, ok := c.data[key]
	c.lock.Unlock()
	if !ok {
		return nil, errcode.ErrCacheKeyNotFound
	}

	return rs, nil
}

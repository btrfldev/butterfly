package store

import (
	"fmt"
	"sync"
)

// Memory Key:Value Storage
type CarbineStore[K, C comparable, V any, S func(K, C) bool] struct {
	mu      sync.RWMutex
	carbine map[K]V
}

// Return CarbineStore(Memory Key:Value Storage)
func NewCarbineStore[K, C comparable, V any, S func(K, C) bool]() *CarbineStore[K, C, V, S] {
	return &CarbineStore[K, C, V, S]{
		carbine: make(map[K]V),
	}
}

func (c *CarbineStore[K, C, V, S]) Put(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.carbine[key] = value
	return nil
}

func (c *CarbineStore[K, C, V, S]) List(search S, comp C) (keys []K, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for key := range c.carbine {
		if search(key, comp) {
			keys = append(keys, key)
		}

	}

	return keys, nil
}

func (c *CarbineStore[K, C, V, S]) Get(key K) (value V, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.carbine[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}
	return value, nil
}

func (c *CarbineStore[K, C, V, S]) Update(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.carbine[key]
	if !ok {
		return fmt.Errorf("the key (%v) does not exists", key)
	}

	c.carbine[key] = value
	return nil
}

func (c *CarbineStore[K, C, V, S]) Delete(key K) (value V, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.carbine[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}

	delete(c.carbine, key)
	return value, nil
}

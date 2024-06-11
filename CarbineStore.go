package butterfly

import (
	"fmt"
	"sync"
)

//Memory Key:Value Storage
type CarbineStore[K comparable, V any] struct {
	mu   sync.RWMutex
	carbine map[K]V
}

//Return DustStroe(Memory Key:Value Storage)
func NewDustStore[K comparable, V any]() *CarbineStore[K, V] {
	return &CarbineStore[K, V]{
		carbine: make(map[K]V),
	}
}


func (c *CarbineStore[K, V]) Put(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.carbine[key] = value
	return nil
}

func (c *CarbineStore[K, V]) List() (keys []K, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for key := range c.carbine {
		keys = append(keys, key)
	}

	return keys, nil
}

func (c *CarbineStore[K, V]) Get(key K) (value V, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.carbine[key]
	if !ok{
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}
	return value, nil
}

func (c *CarbineStore[K, V]) Update(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.carbine[key]
	if !ok {
		return fmt.Errorf("the key (%v) does not exists", key)
	}

	c.carbine[key] = value
	return nil
}

func (c *CarbineStore[K, V]) Delete(key K) (value V, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.carbine[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}

	delete(c.carbine, key)
	return value, nil
}
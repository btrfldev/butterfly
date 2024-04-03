package butterfly

import (
	"sync"
	"fmt"
)

//Memory Key:Value Storage
type DustStore[K comparable, V any] struct {
	mu   sync.RWMutex
	dust map[K]V
}

//Return AirStroe(Memory Key:Value Storage)
func NewDustStore[K comparable, V any]() *DustStore[K, V] {
	return &DustStore[K, V]{
		dust: make(map[K]V),
	}
}


// Has checks if the given is present in the store
// NOTE: DON`T use in CONCURENT safe. Should be used whith the lock.
func (a *DustStore[K, V]) Has(key K) bool {
	_, ok := a.dust[key]
	return ok
}

func (a *DustStore[K, V]) Put(key K, value V) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.dust[key] = value
	return nil
}

func (a *DustStore[K, V]) Get(key K) (value V, err error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	value, ok := a.dust[key]
	if !ok{
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}
	return value, nil
}

func (a *DustStore[K, V]) Update(key K, value V) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.Has(key) {
		return fmt.Errorf("the key (%v) does not exists", key)
	}

	a.dust[key] = value
	return nil
}

func (a *DustStore[K, V]) Delete(key K) (value V, err error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	value, ok := a.dust[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}

	delete(a.dust, key)
	return value, nil
}
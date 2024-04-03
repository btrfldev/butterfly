package butterfly

import (
	"sync"
	"fmt"
)

//Memory Key:Value Storage
type AirStore[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

//Return AirStroe(Memory Key:Value Storage)
func NewAirStore[K comparable, V any]() *AirStore[K, V] {
	return &AirStore[K, V]{
		data: make(map[K]V),
	}
}


// Has checks if the given is present in the store
// NOTE: DON`T use in CONCURENT safe. Should be used whith the lock.
func (a *AirStore[K, V]) Has(key K) bool {
	_, ok := a.data[key]
	return ok
}

func (a *AirStore[K, V]) Put(key K, value V) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.data[key] = value
	return nil
}

func (a *AirStore[K, V]) Get(key K) (value V, err error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	value, ok := a.data[key]
	if !ok{
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}
	return value, nil
}

func (a *AirStore[K, V]) Update(key K, value V) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.Has(key) {
		return fmt.Errorf("the key (%v) does not exists", key)
	}

	a.data[key] = value
	return nil
}

func (a *AirStore[K, V]) Delete(key K) (value V, err error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	value, ok := a.data[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}

	delete(a.data, key)
	return
}
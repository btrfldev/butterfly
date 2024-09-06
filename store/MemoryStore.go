package store

import (
	"fmt"
	"sync"
)

// Memory Key:Value Storage
type MemoryStore struct {
	mu      sync.RWMutex
	memory map[string]string
}


func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		memory: make(map[string]string),
		//mu:     sync.RWMutex{},
	}
}

func (m *MemoryStore) Put(key string, value string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.memory[key] = value
	return nil
}

func (m *MemoryStore) List(search func(k string, c string) bool, comp string) (keys []string, err error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for key := range m.memory {
		if search(key, comp) {
			keys = append(keys, key)
		}

	}

	return keys, nil
}

func (m *MemoryStore) Get(key string) (value string, err error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.memory[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}
	return value, nil
}

func (m *MemoryStore) Update(key string, value string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.memory[key]
	if !ok {
		return fmt.Errorf("the key (%v) does not exists", key)
	}

	m.memory[key] = value
	return nil
}

func (m *MemoryStore) Delete(key string) (value string, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	value, ok := m.memory[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}

	delete(m.memory, key)
	return value, nil
}

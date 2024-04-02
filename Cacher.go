package butterfly

import "sync"

type Casher struct {
	mu   sync.RWMutex
	data map[string][]byte
}

package server

import "sync"

type inMemoryCallCounter struct {
	cache map[string]int
	lck   sync.RWMutex
}

func (c *inMemoryCallCounter) Increment(id string, inc int) error {
	c.lck.Lock()
	defer c.lck.Unlock()
	c.cache[id] += inc
	return nil
}

func (c *inMemoryCallCounter) Lookup(id string) (int, error) {
	c.lck.RLock()
	defer c.lck.RUnlock()
	return c.cache[id], nil
}

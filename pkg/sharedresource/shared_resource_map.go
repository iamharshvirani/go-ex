package sharedresource

import (
	"fmt"
	"sync"
)

// sync.Map for concurrent map access
// This is a specialized map type designed for concurrent access without explicit locking by the user for common operations like Store and Load.

// ConcurrentCache is a simple concurrent-safe cache using sync.Map
type ConcurrentCache struct {
	data sync.Map // sync.Map is safe for concurrent use
}

// Store adds or updates a key-value pair in the cache
func (c *ConcurrentCache) Store(key string, value interface{}) {
	c.data.Store(key, value)
	fmt.Printf("Stored key: %s\n", key)
}

// Load retrieves a value from the cache
func (c *ConcurrentCache) Load(key string) (interface{}, bool) {
	value, ok := c.data.Load(key)
	fmt.Printf("Loaded key: %s, found: %t\n", key, ok)
	return value, ok
}

func RunSharedResourceMap() {
	cache := ConcurrentCache{}
	var wg sync.WaitGroup

	// Goroutines to store data
	keysToStore := []string{"user:1", "product:10", "order:abc"}
	for _, key := range keysToStore {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			cache.Store(k, fmt.Sprintf("value_for_%s", k))
		}(key)
	}

	wg.Wait()
	fmt.Println("Finished storing data.")

	// Goroutines to load data
	keysToLoad := []string{"user:1", "product:20", "order:abc", "nonexistent"}
	for _, key := range keysToLoad {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			cache.Load(k)
		}(key)
	}

	wg.Wait()
	fmt.Println("Finished loading data.")
}

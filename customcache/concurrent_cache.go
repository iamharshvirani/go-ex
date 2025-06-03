package customcache

import (
	"fmt"
	"sync"
	"time" // Added for simulating concurrent access
)

// To make the cache thread-safe, we need to protect the shared map from concurrent access using a synchronization mechanism.

// Cache stores key-value pairs
type ConcurrentCache struct {
	mu   sync.RWMutex // Using RWMutex for better performance with many reads
	data map[string]interface{}
}

// NewConcurrentCache creates a new instance of ConcurrentCache
func NewConcurrentCache() *ConcurrentCache {
	return &ConcurrentCache{
		data: make(map[string]interface{}),
	}
}

// Set adds or updates a key-value pair in the cache
func (c *ConcurrentCache) Set(key string, value interface{}) {
	c.mu.Lock()         // Acquire a write lock
	defer c.mu.Unlock() // Release the write lock
	c.data[key] = value
	fmt.Printf("Cache: Set key '%s'\n", key)
}

// Get retrieves a value from the cache
func (c *ConcurrentCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()         // Acquire a read lock
	defer c.mu.RUnlock() // Release the read lock
	value, found := c.data[key]
	fmt.Printf("Cache: Get key '%s' - Found: %t\n", key, found)
	return value, found
}

// Delete removes a key-value pair from the cache
func (c *ConcurrentCache) Delete(key string) {
	c.mu.Lock()         // Acquire a write lock
	defer c.mu.Unlock() // Release the write lock
	delete(c.data, key)
	fmt.Printf("Cache: Deleted key '%s'\n", key)
}

func RunConcurrentCache() {
	cache := NewConcurrentCache()
	var wg sync.WaitGroup

	// Simulate concurrent access
	keys := []string{"user:1", "product:10", "order:abc", "user:1", "product:10"}
	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			cache.Set(k, fmt.Sprintf("value_for_%s_%d", k, time.Now().UnixNano())) // Set with unique value
			cache.Get(k)
		}(key)
	}

	wg.Wait()
	fmt.Println("Concurrent access simulation finished.")
	fmt.Printf("Final cache size: %d\n", len(cache.data)) // Note: len() on map requires lock
}

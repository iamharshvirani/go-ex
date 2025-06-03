package customcache

import "fmt"

// You have a backend service that frequently reads data from a database. To reduce the load on the database and improve response times for common requests, you decide to implement an in-memory cache.
// Describe how you would design a simple in-memory key-value cache in Go to store data retrieved from the database.

// Cache stores key-value pairs
type SimpleCache struct {
	data map[string]interface{} // Using interface{} to store any type of data
}

// NewSimpleCache creates a new instance of SimpleCache
func NewSimpleCache() *SimpleCache {
	return &SimpleCache{
		data: make(map[string]interface{}),
	}
}

// Set adds or updates a key-value pair in the cache
func (c *SimpleCache) Set(key string, value interface{}) {
	c.data[key] = value
	fmt.Printf("Cache: Set key '%s'\n", key)
}

// Get retrieves a value from the cache
func (c *SimpleCache) Get(key string) (interface{}, bool) {
	value, found := c.data[key]
	fmt.Printf("Cache: Get key '%s' - Found: %t\n", key, found)
	return value, found
}

// Delete removes a key-value pair from the cache
func (c *SimpleCache) Delete(key string) {
	delete(c.data, key)
	fmt.Printf("Cache: Deleted key '%s'\n", key)
}

func RunSimpleCache() {
	cache := NewSimpleCache()

	cache.Set("user:123", map[string]string{"name": "Alice"})
	user, found := cache.Get("user:123")
	if found {
		fmt.Printf("Retrieved from cache: %+v\n", user)
	}

	cache.Delete("user:123")
	_, found = cache.Get("user:123") // Should not be found now
}

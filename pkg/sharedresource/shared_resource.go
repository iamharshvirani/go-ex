package sharedresource

import (
	"fmt"
	"sync"
	"time"
)

// Multiple goroutines need to read from and/or write to a shared data structure like a map or a counter.

// SafeCounter is a counter that can be safely accessed concurrently.
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc increments the counter.
func (c *SafeCounter) Inc() {
	c.mu.Lock()         // Acquire the lock
	defer c.mu.Unlock() // Release the lock when the function exits
	c.count++
	fmt.Printf("Incremented count to %d\n", c.count)
	time.Sleep(time.Millisecond * 10) // Simulate some work while holding the lock
}

// Value returns the current value of the counter.
func (c *SafeCounter) Value() int {
	c.mu.Lock()         // Acquire the lock for reading
	defer c.mu.Unlock() // Release the lock
	return c.count
}

func RunSharedResource() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// Launch multiple goroutines to increment the counter
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter.Value()) // Access the final value safely
}

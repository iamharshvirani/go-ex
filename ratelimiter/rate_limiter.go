package ratelimiter

import (
	"fmt"
	"sync"
	"time"
)

// Design and implement a simple, in-memory rate limiter in Go for API requests based on a user ID. Consider concurrency and potential issues.

// RateLimiterConfig holds the rate limit settings
type RateLimiterConfig struct {
	Limit  int           // Maximum number of requests
	Window time.Duration // Time window for the limit
}

// RateLimiter is an in-memory rate limiter
type RateLimiter struct {
	mu     sync.Mutex
	counts map[string]int
	config RateLimiterConfig
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(config RateLimiterConfig) *RateLimiter {
	rl := &RateLimiter{
		counts: make(map[string]int),
		config: config,
	}
	go rl.cleanupRoutine() // Start the cleanup goroutine
	return rl
}

// Allow checks if a request is allowed for a given userID
func (rl *RateLimiter) Allow(userID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Increment the count for the user
	rl.counts[userID]++

	// Check if the count exceeds the limit
	if rl.counts[userID] > rl.config.Limit {
		return false // Rate limited
	}

	return true // Allowed
}

// cleanupRoutine periodically resets the counts
func (rl *RateLimiter) cleanupRoutine() {
	ticker := time.NewTicker(rl.config.Window)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		// Reset all counts after the window
		rl.counts = make(map[string]int)
		fmt.Println("Rate limiter counts reset.")
		rl.mu.Unlock()
	}
}

func RunRateLimiter() {
	config := RateLimiterConfig{
		Limit:  3,
		Window: time.Second * 5,
	}
	limiter := NewRateLimiter(config)

	// Simulate incoming requests
	requests := []string{"userA", "userB", "userA", "userA", "userB", "userA", "userC"}

	for _, user := range requests {
		if limiter.Allow(user) {
			fmt.Printf("Request for %s allowed.\n", user)
		} else {
			fmt.Printf("Request for %s denied (rate limited).\n", user)
		}
		time.Sleep(time.Millisecond * 500) // Simulate request processing time
	}

	// Let the cleanup routine run for a bit
	time.Sleep(time.Second * 7)
}

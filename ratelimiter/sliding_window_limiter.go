package ratelimiter

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RateLimiterConfig holds the rate limit settings
type SlidingWindowRateLimiterConfig struct {
	Limit         int           // Maximum number of requests
	Window        time.Duration // Sliding time window
	CleanupPeriod time.Duration // How often to run cleanup
}

// RateLimiter implements an in-memory sliding-window rate limiter
type SlidingWindowRateLimiter struct {
	mu sync.RWMutex
	// For each userID, keep a slice of timestamps when requests occurred
	userTimestamps map[string][]time.Time
	config         SlidingWindowRateLimiterConfig
}

// NewRateLimiter creates a new RateLimiter
func NewSlidingWindowRateLimiter(config SlidingWindowRateLimiterConfig) *SlidingWindowRateLimiter {
	rl := &SlidingWindowRateLimiter{
		userTimestamps: make(map[string][]time.Time),
		config:         config,
	}
	go rl.cleanupRoutine()
	return rl
}

// Allow checks if a request is allowed for a given userID
func (rl *SlidingWindowRateLimiter) Allow(userID string) bool {
	now := time.Now()

	// First, lock for writing because we may modify the slice
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Prune old timestamps outside the sliding window
	pruned := rl.pruneOld(userID, now)

	// After pruning, check how many remain
	if len(pruned) >= rl.config.Limit {
		// Already at or above limit
		return false
	}

	// Otherwise, append current timestamp and allow
	rl.userTimestamps[userID] = append(pruned, now)
	return true
}

// pruneOld returns a new slice of timestamps within the window
func (rl *SlidingWindowRateLimiter) pruneOld(userID string, now time.Time) []time.Time {
	windowStart := now.Add(-rl.config.Window)
	oldTimestamps := rl.userTimestamps[userID]
	// Find the first index i where oldTimestamps[i] >= windowStart
	cut := 0
	for cut < len(oldTimestamps) {
		if oldTimestamps[cut].After(windowStart) {
			break
		}
		cut++
	}
	pruned := oldTimestamps[cut:] // all timestamps >= windowStart
	rl.userTimestamps[userID] = pruned
	return pruned
}

// GetRemaining returns how many more requests this user can make right now.
// If they have already hit the limit, returns 0.
func (rl *SlidingWindowRateLimiter) GetRemaining(userID string) int {
	now := time.Now()

	rl.mu.RLock()
	defer rl.mu.RUnlock()

	timestamps := rl.userTimestamps[userID]
	// Find how many are still within window
	windowStart := now.Add(-rl.config.Window)
	count := 0
	for _, ts := range timestamps {
		if ts.After(windowStart) {
			count++
		}
	}
	remaining := rl.config.Limit - count
	if remaining < 0 {
		return 0
	}
	return remaining
}

// GetRetryAfter returns how long until the user can make at least one more request.
// If they are under limit, returns 0.
func (rl *SlidingWindowRateLimiter) GetRetryAfter(userID string) time.Duration {
	now := time.Now()

	rl.mu.RLock()
	defer rl.mu.RUnlock()

	timestamps := rl.userTimestamps[userID]
	if len(timestamps) < rl.config.Limit {
		return 0
	}

	// We know len(timestamps) >= Limit. We need the timestamp at index
	// (len - Limit). That is, once the “oldest of the last Limit” falls out
	// of the window, they can make one more. (Sliding-window logic.)
	// First prune to be safe:
	windowStart := now.Add(-rl.config.Window)
	pruned := timestamps[:0]
	for _, ts := range timestamps {
		if ts.After(windowStart) {
			pruned = append(pruned, ts)
		}
	}
	if len(pruned) < rl.config.Limit {
		// After pruning, they’re under limit
		return 0
	}

	// The “earliest” of those last Limit timestamps:
	earliest := pruned[0] // sorted by insertion time
	retryAfter := earliest.Add(rl.config.Window).Sub(now)
	if retryAfter < 0 {
		return 0
	}
	return retryAfter
}

// cleanupRoutine runs periodically to wipe out empty users and prune old timestamps.
// This prevents unbounded memory growth in long-running processes.
func (rl *SlidingWindowRateLimiter) cleanupRoutine() {
	ticker := time.NewTicker(rl.config.CleanupPeriod)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		windowStart := now.Add(-rl.config.Window)
		for userID, timestamps := range rl.userTimestamps {
			// Prune old timestamps
			pruned := timestamps[:0]
			for _, ts := range timestamps {
				if ts.After(windowStart) {
					pruned = append(pruned, ts)
				}
			}
			if len(pruned) == 0 {
				// No recent requests—remove user entry entirely
				delete(rl.userTimestamps, userID)
			} else {
				rl.userTimestamps[userID] = pruned
			}
		}
		rl.mu.Unlock()
		fmt.Println("[Cleanup] Completed pruning old entries")
	}
}

func RunSlidingWindowRateLimiter() {
	// Example configuration: max 5 requests per 10 seconds, cleanup every 5 seconds
	config := SlidingWindowRateLimiterConfig{
		Limit:         5,
		Window:        10 * time.Second,
		CleanupPeriod: 5 * time.Second,
	}
	limiter := NewSlidingWindowRateLimiter(config)

	// Simulate a bursty workload from multiple users
	var wg sync.WaitGroup
	users := []string{"alice", "bob", "charlie"}

	for _, user := range users {
		// Each user spawns a goroutine that sends a random number of requests
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			for i := 0; i < 12; i++ {
				allowed := limiter.Allow(u)
				if allowed {
					fmt.Printf("[%s] Request %2d: ✅ allowed (remaining=%d)\n",
						u, i+1, limiter.GetRemaining(u))
				} else {
					retry := limiter.GetRetryAfter(u)
					fmt.Printf("[%s] Request %2d: 🚫 rate limited (retry after %v)\n",
						u, i+1, retry.Truncate(time.Millisecond))
				}
				// Random sleep between 0–2s
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			}
		}(user)
	}

	// Also, simulate a tail-end check for a user making a request after the window
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(12 * time.Second)
		user := "alice"
		allowed := limiter.Allow(user)
		fmt.Printf("[Late] %s after 12s: allowed=%v (should be allowed because older entries are expired)\n", user, allowed)
	}()

	wg.Wait()

	// Let the cleanup goroutine run one more time before exit
	time.Sleep(6 * time.Second)
	fmt.Println("Done.")
}

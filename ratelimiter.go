package ratelimiter

import (
	"sync"
	"time"
)

// RateLimiter structure for storing rate limiter data.
type RateLimiter struct {
	rate     int
	bucket   int
	max      int
	interval time.Duration
	mu       sync.Mutex
}

// New created new RateLimiter.
func New(rate, max int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:     rate,
		max:      max,
		bucket:   max,
		interval: interval,
	}
}

// IsAllow checks whether the action can be performed and, if so, reduces the number of available tokens.
func (rl *RateLimiter) IsAllow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.bucket > 0 {
		rl.bucket--
		return true
	}
	return false
}

// refill periodically adds tokens to the jar until it reaches its maximum capacity.
func (rl *RateLimiter) refill() {
	ticker := time.NewTicker(rl.interval)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		if rl.bucket < rl.max {
			rl.bucket += rl.rate
			if rl.bucket > rl.max {
				rl.bucket = rl.max
			}
		}
		rl.mu.Unlock()
	}
}

// Start starts the process of refilling tokens.
func (rl *RateLimiter) Start() {
	go rl.refill()
}

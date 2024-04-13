package ratelimiter

import (
	"sync"
	"time"
)

// RateLimiter structure for storing rate limiter data.
type RateLimiter struct {
	rate       int
	bucket     int
	max        int
	interval   time.Duration
	lastUpdate time.Time
	mu         sync.Mutex
}

// New created new RateLimiter.
func New(rate, max int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		rate:       rate,
		max:        max,
		bucket:     max,
		interval:   interval,
		lastUpdate: time.Now(),
	}
}

// IsAllow checks whether the action can be performed and, if so, reduces the number of available tokens.
func (rl *RateLimiter) IsAllow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastUpdate)

	if elapsed >= rl.interval {
		tokensToAdd := (int(elapsed / rl.interval)) * rl.rate
		if tokensToAdd > 0 {
			rl.bucket += tokensToAdd
			if rl.bucket > rl.max {
				rl.bucket = rl.max
			}
			rl.lastUpdate = now
		}
	}

	if rl.bucket > 0 {
		rl.bucket--
		return true
	}
	return false
}

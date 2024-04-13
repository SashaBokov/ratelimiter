# RateLimiter

A simple, efficient rate limiting library written in Go, designed to control the rate of requests an entity can make within a certain time frame. It uses a token bucket algorithm to allow or deny requests based on the available tokens, providing a straightforward way to implement rate limiting in Go applications.

## Features

- Easy to integrate and use in any Go application.
- Configurable rate limiting options to suit different use cases.
- Supports custom time intervals for more flexible rate control.
- Thread-safe implementation using Go's concurrency primitives.

## Getting Started

### Installation

To start using `ratelimiter`, install Go and run `go get`:

```bash
go get github.com/SashaBokov/ratelimiter
```

## Usage Examples

The `ratelimiter` library can be applied in various scenarios to ensure that the rate of actions does not exceed the defined limits. Here are some practical examples:

### Limiting Messages Per Second

To limit a user from sending more than 5 messages per second:

```go
rl := ratelimiter.New(5, 5, time.Second)

// Simulating message sending
for i := 0; i < 5; i++ {
	if rl.IsAllow() {
		// Message allowed
	}
}

// The next message within the same second will be blocked
if !rl.IsAllow() {
	// Message blocked
}
```

### Limiting Requests Per Minute

To prevent an IP address from making more than 10,000 requests per minute:

```go
rl := ratelimiter.New(10000, 10000, time.Minute)

// Simulating requests
for i := 0; i < 10000; i++ {
	if rl.IsAllow() {
		// Request allowed
	}
}

// The next request within the same minute will be blocked
if !rl.IsAllow() {
	// Request blocked
}
```

### Limiting Transactions Per Day

To allow a user to have no more than 3 failed card transactions per day:

```go
rl := ratelimiter.New(3, 3, 24*time.Hour)

// Simulating transactions
for i := 0; i < 3; i++ {
	if rl.IsAllow() {
		// Transaction allowed
	}
}

// The next transaction within the same day will be blocked
if !rl.IsAllow() {
	// Transaction blocked
}
```

#### These examples demonstrate how to use ratelimiter to control the rate of actions across various use cases efficiently. All these examples can be found in the `ratelimiter_test.go` file for reference.

package ratelimiter

import (
	"testing"
	"time"
)

func TestRateLimiter_MessagesPerSecond(t *testing.T) {
	rl := New(5, 5, 1)
	rl.Start()
	time.Sleep(time.Second) // Let's wait for the initial replenishment
	passed := 0

	for i := 0; i < 5; i++ {
		if rl.IsAllow() {
			passed++
			t.Logf("Message %d allowed", i+1)
		}
	}

	if passed != 5 {
		t.Errorf("Expected to allow 5 messages, but allowed %d", passed)
	} else {
		t.Logf("All 5 messages were successfully allowed")
	}

	// Checking that the next message will be blocked
	if rl.IsAllow() {
		t.Error("Expected to block the 6th message, but it was allowed")
	} else {
		t.Log("6th message was correctly blocked")
	}
}

func TestRateLimiter_RequestsPerMinute(t *testing.T) {
	rl := New(10000, 10000, time.Minute)
	rl.Start()
	time.Sleep(time.Second) // Let's wait for the initial replenishment
	passed := 0

	for i := 0; i < 10000; i++ {
		if rl.IsAllow() {
			passed++
		}
	}

	if passed != 10000 {
		t.Errorf("Expected to allow 10000 requests, but allowed %d", passed)
	} else {
		t.Logf("All 10000 requests were successfully allowed")
	}

	// Checking that the next request will be blocked
	if rl.IsAllow() {
		t.Error("Expected to block the 10001st request, but it was allowed")
	} else {
		t.Log("10001st request was correctly blocked")
	}
}

func TestRateLimiter_TransactionsPerDay(t *testing.T) {
	rl := New(3, 3, time.Hour*24)
	rl.Start()
	time.Sleep(time.Second) // Let's wait for the initial replenishment
	passed := 0

	for i := 0; i < 3; i++ {
		if rl.IsAllow() {
			passed++
			t.Logf("Transaction %d allowed", i+1)
		}
	}

	if passed != 3 {
		t.Errorf("Expected to allow 3 transactions, but allowed %d", passed)
	} else {
		t.Logf("All 3 transactions were successfully allowed")
	}

	// Checking that the next transaction will be blocked
	if rl.IsAllow() {
		t.Error("Expected to block the 4th transaction, but it was allowed")
	} else {
		t.Log("4th transaction was correctly blocked")
	}
}

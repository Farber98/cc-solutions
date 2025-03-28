package main

import (
	"testing"
	"time"

	"github.com/test-go/testify/require"
)

func TestBucketCreation(t *testing.T) {
	tb := NewTokenBucket(10, 1*time.Second, 5*time.Minute)
	ip := "192.168.1.1"
	b := tb.getBucket(ip)

	if b.tokens != 10 {
		t.Errorf("Expected 10 tokens, got %d", b.tokens)
	}
}

func TestTokenRefill(t *testing.T) {
	tb := NewTokenBucket(5, 100*time.Millisecond, 5*time.Minute)
	ip := "192.168.1.1"
	b := tb.getBucket(ip)

	b.mu.Lock()
	b.tokens = 0
	b.mu.Unlock()

	time.Sleep(1 * time.Second)
	b.mu.RLock()
	defer b.mu.RUnlock()
	if b.tokens != tb.maxTokens {
		t.Errorf("tokens should be refilled")
	}
}

func TestAllowTb(t *testing.T) {
	tb := NewTokenBucket(3, 1*time.Second, 5*time.Minute)
	ip := "192.168.1.1"
	for range 3 {
		require.True(t, tb.Allow(ip))
	}
	require.False(t, tb.Allow(ip))
}

func TestCleanup(t *testing.T) {
	tb := NewTokenBucket(3, 100*time.Millisecond, 1*time.Second)
	ip := "192.168.1.1"
	tb.Allow(ip)
	_, exists := tb.ipBuckets[ip]
	require.True(t, exists)
	time.Sleep(2 * time.Second)
	_, exists = tb.ipBuckets[ip]
	require.False(t, exists)
}

func TestGetStateTb(t *testing.T) {
	tb := NewTokenBucket(3, 1*time.Millisecond, 10*time.Second)
	ip := "192.168.1.1"
	mp := tb.GetState(ip)
	require.Equal(t, 3, mp["tokens"])
}

package main

import (
	"log"
	"sync"
	"time"
)

type TokenBucket struct {
	ipBuckets map[string]*bucket
	mu        sync.RWMutex
	// The maximum number of tokens in the bucket
	maxTokens int
	// The rate at which tokens are added to the bucket
	rate time.Duration
	// The rate at which we cleanup unused buckets
	expiration time.Duration
}

func NewTokenBucket(maxTokens int, rate, expiration time.Duration) *TokenBucket {
	tb := &TokenBucket{
		ipBuckets:  make(map[string]*bucket),
		mu:         sync.RWMutex{},
		maxTokens:  maxTokens,
		rate:       rate,
		expiration: expiration,
	}

	go tb.refill()
	go tb.cleanup()
	return tb
}

func (tb *TokenBucket) getBucket(ip string) *bucket {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	if b, ok := tb.ipBuckets[ip]; ok {
		return b
	}

	b := &bucket{tokens: tb.maxTokens}
	tb.ipBuckets[ip] = b
	return b
}

type bucket struct {
	// The number of tokens in the bucket
	tokens   int
	mu       sync.RWMutex
	lastUsed time.Time
}

func (tb *TokenBucket) refill() {
	ticker := time.NewTicker(tb.rate)
	defer ticker.Stop()

	for range ticker.C {
		tb.mu.RLock()
		buckets := make([]*bucket, 0, len(tb.ipBuckets))
		for _, b := range tb.ipBuckets {
			buckets = append(buckets, b)
		}
		tb.mu.RUnlock()

		for _, b := range buckets {
			b.mu.Lock()
			if b.tokens < tb.maxTokens {
				b.tokens++
			}
			b.mu.Unlock()
		}
	}
}

func (tb *TokenBucket) cleanup() {
	ticker := time.NewTicker(tb.expiration)
	defer ticker.Stop()

	for range ticker.C {
		tb.mu.RLock()
		expiredBuckets := make([]string, 0)
		for ip, b := range tb.ipBuckets {
			b.mu.RLock()
			if time.Now().After(b.lastUsed.Add(tb.expiration)) {
				expiredBuckets = append(expiredBuckets, ip)
			}
			b.mu.RUnlock()
		}
		tb.mu.RUnlock()

		tb.mu.RLock()
		for _, ip := range expiredBuckets {
			delete(tb.ipBuckets, ip)
			log.Printf("Deleted bucket for IP: %s due to inactivity", ip)
		}
		tb.mu.RUnlock()
	}
}

func (tb *TokenBucket) Allow(ip string) bool {
	b := tb.getBucket(ip)
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.tokens == 0 {
		return false
	}
	b.tokens--
	b.lastUsed = time.Now()
	return true
}

func (tb *TokenBucket) GetState(ip string) map[string]any {
	b := tb.getBucket(ip)

	b.mu.RLock()
	defer b.mu.RUnlock()
	return map[string]any{
		"tokens":   b.tokens,
		"lastUsed": b.lastUsed,
	}
}

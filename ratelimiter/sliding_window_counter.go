package main

import (
	"sync"
	"time"
)

type SlidingWindowCounter struct {
	threshold           int
	currentWindowCount  int
	previousWindowCount int
	currentWindowStart  time.Time
	windowDuration      time.Duration
	mu                  sync.RWMutex
}

func NewSlidingWindowCounter(threshold int, windowDuration time.Duration) *SlidingWindowCounter {
	swc := &SlidingWindowCounter{
		threshold:          threshold,
		currentWindowStart: time.Now(),
		windowDuration:     windowDuration,
		mu:                 sync.RWMutex{},
	}
	return swc
}

func (swc *SlidingWindowCounter) Allow(ip string) bool {
	swc.mu.Lock()
	defer swc.mu.Unlock()
	// Check if we have to change windows
	elapsed := time.Since(swc.currentWindowStart)
	for elapsed >= swc.windowDuration {
		swc.previousWindowCount = swc.currentWindowCount
		swc.currentWindowCount = 0
		swc.currentWindowStart = swc.currentWindowStart.Add(swc.windowDuration)
		elapsed = time.Since(swc.currentWindowStart)
	}

	// Weighted sum: (1 - elapsedRatio) * previous + current
	elapsedRatio := float64(elapsed) / float64(swc.windowDuration)
	weightedCount := (1-elapsedRatio)*float64(swc.previousWindowCount) + float64(swc.currentWindowCount)

	if weightedCount >= float64(swc.threshold) {
		return false
	}

	swc.currentWindowCount++
	return true
}

func (swc *SlidingWindowCounter) GetState(ip string) map[string]any {
	swc.mu.RLock()
	defer swc.mu.RUnlock()
	return map[string]any{
		"currentCount":       swc.currentWindowCount,
		"previousCount":      swc.previousWindowCount,
		"duration":           swc.windowDuration,
		"currentWindowStart": swc.currentWindowStart,
	}

}

package main

import (
	"sync"
	"time"
)

type FixedWindow struct {
	count int
	size  int
	mu    sync.RWMutex
}

func NewFixedWindow(size int, reset time.Duration) *FixedWindow {
	fw := &FixedWindow{
		size: size,
		mu:   sync.RWMutex{},
	}
	go fw.resetLoop(reset)
	return fw
}

func (fw *FixedWindow) resetLoop(reset time.Duration) {
	ticker := time.NewTicker(reset)
	defer ticker.Stop()

	for range ticker.C {
		fw.mu.Lock()
		fw.count = 0
		fw.mu.Unlock()
	}
}

func (fw *FixedWindow) Allow(ip string) bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	if fw.count == fw.size {
		return false
	}
	fw.mu.RUnlock()

	fw.mu.Lock()
	defer fw.mu.Unlock()
	fw.count++
	return true
}

func (fw *FixedWindow) GetState(ip string) map[string]any {
	fw.mu.RLock()
	defer fw.mu.RUnlock()
	return map[string]any{
		"count": fw.count,
		"size":  fw.size,
	}
}

package main

import (
	"sync"
	"time"
)

type SlidingWindowLog struct {
	threshold int
	retention time.Duration
	ipToLogs  map[string][]time.Time
	mu        sync.RWMutex
}

func NewSlidingWindowLog(threshold int, retention time.Duration) *SlidingWindowLog {
	return &SlidingWindowLog{
		threshold: threshold,
		retention: retention,
		ipToLogs:  make(map[string][]time.Time),
	}
}

func (sl *SlidingWindowLog) Allow(ip string) bool {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	if _, exists := sl.ipToLogs[ip]; !exists {
		sl.ipToLogs[ip] = []time.Time{time.Now()}
		return len(sl.ipToLogs[ip]) <= sl.threshold
	}

	tsArr := sl.ipToLogs[ip]
	var expIdx int
	for i := len(tsArr) - 1; i > 0; i-- {
		if time.Since(tsArr[i]) > sl.retention {
			expIdx = i
			break
		}
	}

	tsArr = tsArr[expIdx:]
	if len(tsArr) == sl.threshold {
		return false
	}

	tsArr = append(tsArr, time.Now())
	sl.ipToLogs[ip] = tsArr
	return true
}

func (sl *SlidingWindowLog) GetState(ip string) map[string]any {
	sl.mu.RLock()
	defer sl.mu.RUnlock()

	return map[string]any{
		"count":     len(sl.ipToLogs[ip]),
		"threshold": sl.threshold,
		"reset":     sl.retention,
	}
}

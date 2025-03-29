package main

import (
	"testing"
	"time"

	"github.com/test-go/testify/require"
)

func TestSlidingWindowCounterCreation(t *testing.T) {
	swc := NewSlidingWindowCounter(10, 10*time.Second)
	require.Equal(t, swc.currentWindowCount, 0)
	require.Equal(t, swc.threshold, 10)
	require.Equal(t, swc.previousWindowCount, 0)
	require.Equal(t, swc.windowDuration, 10*time.Second)
	require.NotEmpty(t, swc.currentWindowStart)
}

func TestAllowSwc(t *testing.T) {
	swc := NewSlidingWindowCounter(10, 1*time.Second)
	ip := "192.168.1.1"
	time.Sleep(900 * time.Millisecond)
	for range 10 {
		require.True(t, swc.Allow(ip))
	}
	time.Sleep(100 * time.Millisecond)
	require.True(t, swc.Allow(ip))
	require.False(t, swc.Allow(ip))
}

func TestGetStateSwc(t *testing.T) {
	swc := NewSlidingWindowCounter(5, 2*time.Second) // threshold=5, window=2s
	ip := "10.0.0.2"

	require.True(t, swc.Allow(ip))
	require.True(t, swc.Allow(ip))

	state := swc.GetState(ip)
	require.Equal(t, 2, state["currentCount"], "currentCount should be 2 after 2 requests")
	require.Equal(t, 0, state["previousCount"], "previousCount should still be 0")
}

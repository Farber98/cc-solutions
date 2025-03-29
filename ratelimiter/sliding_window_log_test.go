package main

import (
	"testing"
	"time"

	"github.com/test-go/testify/require"
)

func TestSlidingWindowLogCreation(t *testing.T) {
	slw := NewSlidingWindowLog(10, 10*time.Second)
	require.Equal(t, slw.retention, 10*time.Second)
	require.Equal(t, slw.threshold, 10)
}

func TestAllowSwl(t *testing.T) {
	slw := NewSlidingWindowLog(5, 10*time.Second)
	ip := "192.168.1.1"

	for range 5 {
		require.True(t, slw.Allow(ip))
	}
	require.False(t, slw.Allow(ip))
}

func TestGetStateSwl(t *testing.T) {
	slw := NewSlidingWindowLog(5, 10*time.Second)
	ip := "192.168.1.1"

	for range 5 {
		require.True(t, slw.Allow(ip))
	}

	state := slw.GetState(ip)
	require.Equal(t, state["count"], 5)
	require.Equal(t, state["threshold"], slw.threshold)
	require.Equal(t, state["reset"], slw.retention)
}

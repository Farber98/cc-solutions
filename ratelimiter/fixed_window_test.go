package main

import (
	"testing"
	"time"

	"github.com/test-go/testify/require"
)

func TestFixedWindowCreation(t *testing.T) {
	fw := NewFixedWindow(10, 1*time.Minute)
	require.Equal(t, fw.count, 0)
	require.Equal(t, fw.size, 10)
}

func TestAllowFw(t *testing.T) {
	fw := NewFixedWindow(10, 1*time.Minute)
	for range 10 {
		require.True(t, fw.Allow(""))
	}
	require.False(t, fw.Allow(""))
}

func TestGetStateFw(t *testing.T) {
	fw := NewFixedWindow(10, 1*time.Minute)
	fw.Allow("")
	mp := fw.GetState("")
	require.Equal(t, mp["size"], 10)
	require.Equal(t, mp["count"], 1)
}

func TestResetLoop(t *testing.T) {
	fw := NewFixedWindow(10, 1*time.Second)
	for range 10 {
		fw.Allow("")
	}
	time.Sleep(1 * time.Second)
	require.Equal(t, fw.count, 0)
}

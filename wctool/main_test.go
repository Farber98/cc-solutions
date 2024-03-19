package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/Farber98/cc-solutions/wctool/cli"
)

func TestMain_InvalidCommand(t *testing.T) {
	// Save original args and restore them at the end of the test
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// Set os.Args to include an invalid command
	os.Args = []string{"main", "invalid"}

	// Use a buffer to capture the output instead of using os.Stdout
	var buf bytes.Buffer
	err := cli.ExecuteCommand("invalid", &buf)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Check if error message matches the expected message
	expectedErrorMessage := "command invalid not found"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message %q, got %q", expectedErrorMessage, err.Error())
	}
}

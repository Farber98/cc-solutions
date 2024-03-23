package commands

import (
	"bytes"
	"os"
	"testing"

	"github.com/Farber98/cc-solutions/compress/cli"
	"github.com/Farber98/cc-solutions/compress/file"
)

func TestCmdCount_Success(t *testing.T) {
	data := []byte("abbcaabbccc") // 3 different chars.

	f := &file.DefaultFile{}

	// Create a temporary test file with data
	filePath, cleanup, err := f.CreateTempFileWithData(data)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	// Set os.Args to include the -count flag and the temporary file path
	os.Args = []string{"", "-count", filePath}

	// Use a buffer to capture the output instead of using os.Stdout
	var buf bytes.Buffer
	err = cli.ExecuteCommand("-count", &buf)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestCmdCount_NoFilePathProvided(t *testing.T) {
	// Set os.Args to include the -count flag
	os.Args = []string{"", "-count"}

	// Use a buffer to capture the output instead of using os.Stdout
	var buf bytes.Buffer
	err := cli.ExecuteCommand("-count", &buf)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Check if error message matches the expected message
	expectedErrorMessage := "usage: go run main.go -count [filePath]"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message %q, got %q", expectedErrorMessage, err.Error())
	}
}

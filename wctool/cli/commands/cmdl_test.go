package commands

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/Farber98/cc-solutions/wctool/cli"
	"github.com/Farber98/cc-solutions/wctool/file"
)

func TestCmdL_Success(t *testing.T) {
	data := []byte(`ja 
	jaja 
	jajaja`) // 3 new lines.

	// Create a temporary test file with data
	filePath, cleanup, err := file.CreateTempFileWithData(data)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	// Set os.Args to include the -c flag and the temporary file path
	os.Args = []string{"", "-l", filePath}

	// Use a buffer to capture the output instead of using os.Stdout
	var buf bytes.Buffer
	err = cli.ExecuteCommand("-l", &buf)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if output matches expected. 3 lines.
	expectedOutput := fmt.Sprintf("%d %s\n", 3, filePath)
	if got := buf.String(); got != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, got)
	}
}

func TestCmdL_NoFilePathProvided(t *testing.T) {
	// Set os.Args to include the -c flag
	os.Args = []string{"", "-l"}

	// Use a buffer to capture the output instead of using os.Stdout
	var buf bytes.Buffer
	err := cli.ExecuteCommand("-l", &buf)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Check if error message matches the expected message
	expectedErrorMessage := "usage: go run main.go -l [filePath]"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message %q, got %q", expectedErrorMessage, err.Error())
	}
}

func TestCmdL_FileNotFound(t *testing.T) {
	// Set os.Args to include the -c flag and a non-existing file path
	os.Args = []string{"", "-l", "non-existing-file"}

	// Use a buffer to capture the output instead of using os.Stdout
	var buf bytes.Buffer
	err := cli.ExecuteCommand("-l", &buf)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Check if error message matches the expected message
	expectedErrorMessage := "file not found: non-existing-file"
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message %q, got %q", expectedErrorMessage, err.Error())
	}
}

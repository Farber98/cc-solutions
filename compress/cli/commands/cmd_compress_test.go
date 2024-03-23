package commands

import (
	"bytes"
	"os"
	"testing"

	"github.com/Farber98/cc-solutions/compress/file"
)

func TestCmdCompress_Execute(t *testing.T) {
	// Create a temporary test file with data
	data := []byte("abbcaabbccc")

	// Create a temporary test file with data
	f := &file.DefaultFile{}
	filePath, cleanup, err := f.CreateTempFileWithData(data)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	// Set os.Args to include the compress flag and the temporary file path
	os.Args = []string{"", "compress", filePath}

	// Create an instance of CmdCompress
	cmd := &CmdCompress{}

	// Execute the command
	var buf bytes.Buffer
	err = cmd.Execute(&buf)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

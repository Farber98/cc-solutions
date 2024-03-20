package file

import (
	"bytes"
	"os"
	"testing"
)

func TestReadFileContents(t *testing.T) {
	data := []byte("ja jaja jajaja")

	// Create an instance of DefaultFile
	f := &DefaultFile{}

	// Create a temporary test file with data
	filePath, cleanup, err := createTempFileWithData(data)
	if err != nil {
		t.Fatal(err)
	}

	// ensures the temporary test file is closed and removed
	defer cleanup()

	f.Path = filePath

	// Call readFileContents with the temporary file
	fileContents, err := f.ReadFileContents()
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Check if the read content matches the expected content
	if !bytes.Equal(fileContents, data) {
		t.Errorf("Expected %q, got %q", string(data), string(fileContents))
	}
}

// CreateTempFileWithData creates a temporary file with the given data for testing purposes.
func createTempFileWithData(data []byte) (string, func(), error) {
	// Create a temporary test file
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		return "", nil, err
	}
	filePath := tmpfile.Name()

	// Defer the file close and remove operations
	cleanup := func() {
		tmpfile.Close()
		os.Remove(filePath)
	}

	// Write data to the file
	if _, err := tmpfile.Write(data); err != nil {
		cleanup()
		return "", nil, err
	}

	// Return the file path and a cleanup function
	return filePath, cleanup, nil
}

package file

import (
	"bytes"
	"testing"
)

func TestReadFileContents(t *testing.T) {
	data := []byte("ja jaja jajaja")

	// Create an instance of DefaultFile
	f := &DefaultFile{}

	// Create a temporary test file with data
	filePath, cleanup, err := CreateTempFileWithData(data)
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

func TestCountLines(t *testing.T) {
	tests := []struct {
		name          string
		fileContents  []byte
		expectedLines int
	}{
		{"Empty File", []byte{}, 0},
		{"Single Line", []byte("hello\n"), 1},
		{"Multiple Lines", []byte("hello\nworld\n"), 2},
		{"No Newline at End", []byte("hello\nworld"), 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &DefaultFile{}
			lines := f.CountLines(tt.fileContents)
			if lines != tt.expectedLines {
				t.Errorf("Expected %d lines, got %d", tt.expectedLines, lines)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
	}{
		{
			name:     "Empty input",
			input:    []byte{},
			expected: 0,
		},
		{
			name:     "Single word",
			input:    []byte("word"),
			expected: 1,
		},
		{
			name:     "Multiple words",
			input:    []byte("multiple words in this sentence"),
			expected: 5,
		},
		{
			name:     "Whitespace only",
			input:    []byte("   "),
			expected: 0,
		},
		{
			name:     "Mixed whitespace and words",
			input:    []byte("  word1  word2\tword3\n"),
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &DefaultFile{}
			count := f.CountWords(tt.input)
			if count != tt.expected {
				t.Errorf("Expected %d words, got %d", tt.expected, count)
			}
		})
	}
}

func TestCountCharacters(t *testing.T) {
	data := []byte("¡Hola, mundo!")

	f := &DefaultFile{}

	// Call CountCharacters with the test data
	charCount := f.CountCharacters(data)

	// Check if the character count matches the expected count
	expectedCount := 13
	if charCount != expectedCount {
		t.Errorf("Expected character count %d, got %d", expectedCount, charCount)
	}
}

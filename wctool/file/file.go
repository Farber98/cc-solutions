package file

import (
	"bytes"
	"os"
	"unicode"
)

// File defines the interface for file operations.
type File interface {
	ReadFileContents() ([]byte, error)
	CountLines(fileContents []byte) int
	CountWords(fileContents []byte) int
	CountCharacters(fileContents []byte) int
}

// DefaultFile implements the File interface with default file operations.
type DefaultFile struct {
	Path string
}

// ReadFileContents reads the contents of a file and returns them as a byte slice
func (f *DefaultFile) ReadFileContents() ([]byte, error) {
	// Open the file
	file, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file contents
	fileContents := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil && n == 0 {
			break
		}
		fileContents = append(fileContents, buffer[:n]...)
		if err != nil {
			break
		}
	}

	return fileContents, nil
}

// CountLines counts the number of lines in a byte slice.
func (f *DefaultFile) CountLines(fileContents []byte) int {
	count := 0
	for _, b := range fileContents {
		if b == '\n' {
			count++
		}
	}
	// Add 1 to count the last line if it doesn't end with a newline character
	if len(fileContents) > 0 && fileContents[len(fileContents)-1] != '\n' {
		count++
	}
	return count
}

// CountWords counts the number of words in a byte slice.
func (f *DefaultFile) CountWords(fileContents []byte) int {
	wordCount := 0
	inWord := false

	for _, b := range fileContents {
		// At each byte, check if it is a whitespace character (space, tab, newline, etc.).
		if unicode.IsSpace(rune(b)) {
			// If it encounters a whitespace character and was previously in a word, it increments the word count.
			if inWord {
				wordCount++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	// Handles the case where the last word in the file is not followed by a whitespace character.
	if inWord {
		wordCount++
	}

	return wordCount
}

// CountCharacters counts the number of characters in a byte slice.
func (f *DefaultFile) CountCharacters(fileContents []byte) int {
	return len(bytes.Runes(fileContents))
}

// CreateTempFileWithData creates a temporary file with the given data.
func CreateTempFileWithData(data []byte) (string, func(), error) {
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

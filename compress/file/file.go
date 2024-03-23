package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Constants for header markers
const (
	HeaderStart = "HS"
	HeaderEnd   = "HE"
)

// File defines the interface for file operations.
type File interface {
	ReadFileContents() ([]byte, error)
	WriteCodeTableOnHeader(codesTable map[byte]string, outputPath string) error
	WriteTextAfterHeader(byteText []byte, outputPath string) error
	CreateTempFileWithData(data []byte) (string, func(), error)
	ReadReverseLookupCodesTableFromHeader(filePath string) (map[string]byte, error)
	CreateNewFile(fileName string) (*os.File, error)
	ReadTextAfterHeader(filePath string) ([]byte, error)
}

// DefaultFile implements the File interface with default file operations.
type DefaultFile struct{}

// ReadFileContents reads the contents of a file and returns them as a byte slice
func (f *DefaultFile) ReadFileContents(path string) ([]byte, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file contents
	var fileContents []byte
	buffer := make([]byte, 4096) // Increased buffer size
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

// WriteCodeTableOnHeader writes the header section to the output file
func (f *DefaultFile) WriteCodeTableOnHeader(codesTable map[byte]string, outputPath string) error {
	// Open the output file for writing
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// Write the header start marker
	_, err = writer.WriteString(HeaderStart + "\n")
	if err != nil {
		return err
	}

	// Write the size of table.
	_, err = writer.WriteString(fmt.Sprintf("%d\n", len(codesTable)))
	if err != nil {
		return err
	}

	// Write the codeTable as a reverse lookup that will be useful to decompress.
	for char, code := range codesTable {
		_, err := writer.WriteString(fmt.Sprintf("%s,%c\n", code, char))
		if err != nil {
			return err
		}
	}

	// Write the header end marker
	_, err = writer.WriteString(HeaderEnd + "\n")
	if err != nil {
		return err
	}

	// Flush the buffer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

// WriteEncodedText writes the encoded text to the output file after the header.
func (f *DefaultFile) WriteTextAfterHeader(byteText []byte, outputPath string) error {
	// Open the output file for writing, append to the file
	file, err := os.OpenFile(outputPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// Write the encoded text to the file
	_, err = writer.Write(byteText)
	if err != nil {
		return err
	}

	// Flush the buffer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

// createTempFileWithData creates a temporary file with the given data for testing purposes.
func (f *DefaultFile) CreateTempFileWithData(data []byte) (string, func(), error) {
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

// ReadReverseLookupCodesTableFromHeader reads the reversed codesTable from the header of an encoded file.
func (f *DefaultFile) ReadReverseLookupCodesTableFromHeader(filePath string) (map[string]byte, error) {
	// Read the file contents
	contents, err := f.ReadFileContents(filePath)
	if err != nil {
		return nil, err
	}

	// Split the contents by newline to get individual lines
	lines := strings.Split(string(contents), "\n")

	// Check for the header start marker
	if len(lines) == 0 || lines[0] != HeaderStart {
		return nil, fmt.Errorf("invalid header start marker")
	}

	// Parse the number of characters in the code table
	if len(lines) < 2 {
		return nil, fmt.Errorf("missing number of characters in the code table")
	}

	numCodes, err := strconv.Atoi(lines[1])
	if err != nil {
		return nil, fmt.Errorf("invalid number of codes in the code table: %w", err)
	}

	// Parse the code table
	reversedCodesTable := make(map[string]byte, numCodes)
	for i := 2; i < len(lines)-2; i++ { // Start from 2 to skip the start marker and the number of characters, and go to the second to last line (before header end)
		// Separate each character and code
		fields := strings.Split(lines[i], ",")
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid code table format")
		}
		char := fields[1][0]
		code := fields[0]
		// Reverse the table for efficient lookup on decompression.
		reversedCodesTable[code] = byte(char)
	}

	// Check for the header end marker
	if len(lines) < numCodes+3 || lines[numCodes+2] != HeaderEnd {
		return nil, fmt.Errorf("invalid header end marker")
	}

	return reversedCodesTable, nil
}

// CreateNewFile creates a new file with the given file name.
func (f *DefaultFile) CreateNewFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// ReadTextAfterHeader reads the encoded content of a file after the header.
func (f *DefaultFile) ReadTextAfterHeader(filePath string) ([]byte, error) {
	// Read the file contents
	contents, err := f.ReadFileContents(filePath)
	if err != nil {
		return nil, err
	}

	// Split the contents by newline to get individual lines
	lines := strings.Split(string(contents), "\n")

	// Find the index of the line after the header end marker
	startIdx := -1
	for i, line := range lines {
		if line == HeaderEnd {
			startIdx = i + 1
			break
		}
	}

	// Check if the header end marker was found
	if startIdx == -1 {
		return nil, fmt.Errorf("header end marker not found")
	}

	// Concatenate the remaining lines to get the encoded content
	encodedContent := strings.Join(lines[startIdx:], "")
	return []byte(encodedContent), nil
}

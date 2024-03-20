package file

import "os"

// File defines the interface for file operations.
type File interface {
	ReadFileContents() ([]byte, error)
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

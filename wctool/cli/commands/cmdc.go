package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/Farber98/cc-solutions/wctool/file"
)

// CmdC implements the Command interface for the -c command.
type CmdC struct{}

// Execute runs the -c command.
func (c *CmdC) Execute(out io.Writer) error {
	// Check if file name was provided
	if len(os.Args) < 3 {
		return fmt.Errorf("usage: go run main.go -c [filePath]")
	}

	filePath := os.Args[2]

	// If file is not found, return an error.
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", filePath)
	}

	// Create an instance of DefaultFile
	f := &file.DefaultFile{Path: filePath}

	// Read file contents
	fileContents, err := f.ReadFileContents()
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Print file byte count of given file name
	fmt.Fprintf(out, "%d %s\n", len(fileContents), filePath)
	return nil
}

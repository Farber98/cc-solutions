package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/Farber98/cc-solutions/wctool/file"
)

// CmdAll implements the Command interface for the -all command.
type CmdAll struct{}

// Execute runs the -all command.
func (c *CmdAll) Execute(out io.Writer) error {
	// Check if file name was provided
	if len(os.Args) < 3 {
		return fmt.Errorf("usage: go run main.go -all [filePath]")
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

	// Print all commands info
	fmt.Fprintf(out, "%d %d %d %d %s\n",
		len(fileContents),               // -c bytes
		f.CountLines(fileContents),      // -l lines
		f.CountWords(fileContents),      // -w words
		f.CountCharacters(fileContents), // -m characters
		filePath)

	return nil
}

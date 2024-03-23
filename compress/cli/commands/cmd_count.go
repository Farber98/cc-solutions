package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/Farber98/cc-solutions/compress/file"
	"github.com/Farber98/cc-solutions/compress/frequency"
)

// CmdCount implements the Command interface for the -count command.
type CmdCount struct{}

// Execute runs the -count command.
func (c *CmdCount) Execute(out io.Writer) error {
	// Check if file name was provided
	if len(os.Args) < 3 {
		return fmt.Errorf("usage: go run main.go -count [filePath]")
	}

	filePath := os.Args[2]

	// Create an instance of DefaultFile
	f := &file.DefaultFile{}

	// Read file contents
	contents, err := f.ReadFileContents(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Calculate character frequencies
	calculator := &frequency.DefaultCalculator{}

	frequencies := calculator.CalculateFrequencies(contents)

	// Print frequencies
	for char, freq := range frequencies {
		fmt.Printf("Character: %c, Frequency: %d\n", char, freq)
	}

	return nil
}

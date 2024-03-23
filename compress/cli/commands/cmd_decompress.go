package commands

import (
	"fmt"
	"io"
	"os"

	compress "github.com/Farber98/cc-solutions/compress/compression"
	"github.com/Farber98/cc-solutions/compress/file"
)

// CmdDecompress implements the Command interface for the -decompress command.
type CmdDecompress struct{}

// Execute runs the -decompress command.
func (c *CmdDecompress) Execute(out io.Writer) error {
	// Check if file name was provided
	if len(os.Args) < 3 {
		return fmt.Errorf("usage: go run main.go decompress [filePath]")
	}

	filePath := os.Args[2]

	// Create an instance of DefaultFile
	f := &file.DefaultFile{}

	// Read reversed look up table from header
	reverseLookupCodeTable, err := f.ReadReverseLookupCodesTableFromHeader(filePath)
	if err != nil {
		return fmt.Errorf("error reading header: %w", err)
	}

	// Read encoded content after header
	encodedText, err := f.ReadTextAfterHeader(filePath)
	if err != nil {
		return fmt.Errorf("error reading encoded text: %w", err)
	}

	// Decode the encoded text using the reverse lookup table
	decompressor := &compress.DefaultDecompressor{}
	decodedText, err := decompressor.Decode(encodedText, reverseLookupCodeTable)
	if err != nil {
		return fmt.Errorf("error decoding text: %w", err)
	}

	// Write the decoded text to a new file
	outputPath := filePath + ".decompressed"
	newFile, err := f.CreateNewFile(outputPath)
	if err != nil {
		return fmt.Errorf("error creating decompressed file: %w", err)
	}
	defer newFile.Close() // Ensure the file is closed after writing

	err = f.WriteTextAfterHeader(decodedText, outputPath)
	if err != nil {
		return fmt.Errorf("error writing decoded text: %w", err)
	}

	return nil
}

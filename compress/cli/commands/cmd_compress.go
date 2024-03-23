package commands

import (
	"fmt"
	"io"
	"os"

	compress "github.com/Farber98/cc-solutions/compress/compression"
	"github.com/Farber98/cc-solutions/compress/file"
	"github.com/Farber98/cc-solutions/compress/frequency"
	"github.com/Farber98/cc-solutions/compress/huffman"
)

// CmdCompress implements the Command interface for the -compress command.
type CmdCompress struct{}

// Execute runs the -count command.
func (c *CmdCompress) Execute(out io.Writer) error {
	// Check if file name was provided
	if len(os.Args) < 3 {
		return fmt.Errorf("usage: go run main.go compress [filePath]")
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

	// Construct Huffman tree
	h := &huffman.DefaultHuffmanCoding{}
	huffmanTree := h.BuildHuffmanTree(frequencies)

	// Assign codes
	codeTable := make(map[byte]string)
	h.AssignCodes(huffmanTree, "", codeTable)

	// Write header with codeTable
	err = f.WriteCodeTableOnHeader(codeTable, filePath+".compressed")
	if err != nil {
		return fmt.Errorf("error writing header: %w", err)
	}

	// Encode and write compressed data
	compressor := &compress.DefaultCompressor{}
	encodedText := compressor.Encode(contents, codeTable)
	err = f.WriteTextAfterHeader(encodedText, filePath+".compressed")
	if err != nil {
		return fmt.Errorf("error writing encoded text: %w", err)
	}

	fmt.Printf(filePath + ".compressed")
	return nil
}

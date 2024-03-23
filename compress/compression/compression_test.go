package compress

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	// Define the code table
	codes := map[byte]string{'a': "00", 'b': "01", 'c': "10", 'd': "11"}

	// Define the source text
	sourceText := []byte("abcd")

	// Instantiate compressor
	compressor := &DefaultCompressor{}

	// Encode the source text
	encodedBytes := compressor.Encode(sourceText, codes)

	// Convert the expected encoded bytes to bytes
	expectedBytes := []byte{0b00011011}

	// Compare the encoded bytes
	if !bytes.Equal(encodedBytes, expectedBytes) {
		t.Errorf("Expected encoded bytes %v, got %v", expectedBytes, encodedBytes)
	}
}

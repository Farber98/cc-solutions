package compress

import (
	"bytes"
)

// Compressor defines the interface for compression operations.
type Compressor interface {
	Encode(sourceText []byte, codes map[byte]string) []byte
}

// DefaultCompressor implements the Compressor interface with the default compression operations.
type DefaultCompressor struct{}

// Encode encodes the source text using the code table and returns the encoded bytes.
func (c *DefaultCompressor) Encode(sourceText []byte, codes map[byte]string) []byte {
	var buffer bytes.Buffer

	// Encode the source text using the code table
	for _, char := range sourceText {
		buffer.WriteString(codes[char])
	}
	encodedText := buffer.String()

	// Pack bits into bytes
	var byteBuffer []byte
	var byteValue byte
	byteIndex := 0
	for _, bit := range encodedText {
		if byteIndex == 0 {
			byteValue = 0
		}
		if bit == '1' {
			byteValue |= 1 << (7 - byteIndex)
		}
		byteIndex++
		if byteIndex == 8 {
			byteIndex = 0
			byteBuffer = append(byteBuffer, byteValue)
		}
	}

	// Handle remaining bits
	if byteIndex != 0 {
		// Calculate the number of bits to shift left to fill the byte
		shift := 8 - byteIndex
		byteBuffer = append(byteBuffer, byteValue<<shift)
	}

	return byteBuffer
}

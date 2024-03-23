package compress

import "fmt"

// Compressor defines the interface for compression operations.
type Decompressor interface {
	Decode(encodedText []byte, codeTable map[string]byte) ([]byte, error)
}

// DefaultCompressor implements the Compressor interface with the default compression operations.
type DefaultDecompressor struct{}

// Decode decodes the encoded text using the reverse lookup code table and returns the original text.
func (d *DefaultDecompressor) Decode(encodedText []byte, codeTable map[string]byte) ([]byte, error) {
	var decodedText []byte
	currentCode := ""

	for _, byteValue := range encodedText {
		// Convert the byte to a binary string representation
		binaryRep := fmt.Sprintf("%08b", byteValue)
		for _, bit := range binaryRep {
			currentCode += string(bit)
			if char, ok := codeTable[currentCode]; ok {
				decodedText = append(decodedText, char)
				currentCode = ""
			}
		}
	}

	if currentCode != "" {
		return nil, fmt.Errorf("invalid code: %s", currentCode)
	}

	return decodedText, nil
}

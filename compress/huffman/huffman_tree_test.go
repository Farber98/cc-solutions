package huffman

import (
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
	// Define character frequencies
	frequencies := map[byte]int{
		'a': 8,
		'b': 3,
		'c': 1,
		'd': 5,
		'e': 12,
		'f': 6,
	}

	// Instance huffman
	h := &DefaultHuffmanCoding{}

	// Build the Huffman tree
	root := h.BuildHuffmanTree(frequencies)

	// Verify the root frequency
	expectedFreq := 35 // 8 + 3 + 1 + 5 + 12 + 6
	if root.Freq != expectedFreq {
		t.Errorf("Expected root frequency %d, got %d", expectedFreq, root.Freq)
	}

	// Verify the structure of the tree (not the complete tree, but the basic structure)
	if root.Left == nil || root.Right == nil {
		t.Error("Expected both left and right children for the root node")
	}
}

func TestAssignCodes(t *testing.T) {
	// Define character frequencies
	frequencies := map[byte]int{
		'a': 8,
		'b': 3,
		'c': 1,
		'd': 5,
		'e': 12,
		'f': 6,
	}

	// Instance huffman
	h := &DefaultHuffmanCoding{}

	// Build the Huffman tree
	root := h.BuildHuffmanTree(frequencies)

	// Initialize a map to store the binary codes
	codes := make(map[byte]string)

	// Assign codes to each character in the Huffman tree
	h.AssignCodes(root, "", codes)

	// Expected binary codes for characters 'a', 'b', 'c', 'd', 'e', and 'f'
	// For visual guidance: https://cmps-people.ok.ubc.ca/ylucet/DS/Huffman.html
	expected := map[byte]string{
		'a': "01",
		'b': "1001",
		'c': "1000",
		'd': "101",
		'e': "11",
		'f': "00",
	}

	// Check if the assigned codes match the expected codes
	for char, expectedCode := range expected {
		if codes[char] != expectedCode {
			t.Errorf("Expected code %q for character %c, got %q", expectedCode, char, codes[char])
		}
	}
}

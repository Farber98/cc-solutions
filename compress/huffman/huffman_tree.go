package huffman

import (
	"container/heap"

	"github.com/Farber98/cc-solutions/compress/priority_queue"
)

// FrequencyCalculator defines the interface for Huffman coding.
type HuffmanCoding interface {
	BuildHuffmanTree(frequencies map[byte]int) *priority_queue.Node
	AssignCodes(node *priority_queue.Node, code string, codes map[byte]string)
}

// DefaultCalculator implements the Calculator interface with default frequency calculation.
type DefaultHuffmanCoding struct{}

// BuildHuffmanTree builds a Huffman tree from the given character frequencies.
func (h *DefaultHuffmanCoding) BuildHuffmanTree(frequencies map[byte]int) *priority_queue.Node {
	// Populate priority queue with nodes for each character frequency
	pq := priority_queue.NewPriorityQueue(frequencies)

	// Build Huffman tree by merging nodes until we only have one node.
	for pq.Len() > 1 {
		// Remove two nodes with the lowest frequency
		minimum := heap.Pop(&pq).(*priority_queue.Node)
		nextMinimum := heap.Pop(&pq).(*priority_queue.Node)

		// Create a new node with the sum of the frequencies
		merged := &priority_queue.Node{
			Char:     0, // Internal node, not a character
			Freq:     minimum.Freq + nextMinimum.Freq,
			Priority: minimum.Freq + nextMinimum.Freq,
			Left:     minimum,
			Right:    nextMinimum,
		}

		// Push the merged node back into the priority queue
		heap.Push(&pq, merged)
	}

	// Return the root of the Huffman tree
	return heap.Pop(&pq).(*priority_queue.Node)
}

// AssignCodes traverses the Huffman tree to assign binary codes to each character.
func (h *DefaultHuffmanCoding) AssignCodes(node *priority_queue.Node, code string, codes map[byte]string) {
	// Base case: if the node is a leaf (character node), assign the code to the character
	if node.Left == nil && node.Right == nil {
		codes[node.Char] = code
		return
	}

	// Recursive call for the left child with code appended by "0"
	if node.Left != nil {
		h.AssignCodes(node.Left, code+"0", codes)
	}

	// Recursive call for the right child with code appended by "1"
	if node.Right != nil {
		h.AssignCodes(node.Right, code+"1", codes)
	}
}

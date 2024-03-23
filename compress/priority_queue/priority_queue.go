// priority_queue.go

package priority_queue

import "container/heap"

// Node represents a node in the Huffman tree with character, frequency, and priority.
type Node struct {
	Char     byte  // Character
	Freq     int   // Frequency
	Priority int   // Priority. Less freq has higher priority.
	Index    int   // Index in the heap
	Left     *Node // Left child
	Right    *Node // Right child
}

// PriorityQueue implements a priority queue for Huffman tree nodes.
type PriorityQueue []*Node

// Len returns the length of the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two nodes by priority.
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, so we use less than here.
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps two nodes in the priority queue.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds a node to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	node.Index = len(*pq)
	*pq = append(*pq, node)
}

// Pop removes and returns the top node from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.Index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// NewPriorityQueue creates a new priority queue initialized with nodes for each character frequency.
func NewPriorityQueue(frequencies map[byte]int) PriorityQueue {
	pq := make(PriorityQueue, 0)

	for char, freq := range frequencies {
		node := &Node{
			Char:     char,
			Freq:     freq,
			Priority: freq,
		}
		heap.Push(&pq, node)
	}

	heap.Init(&pq)
	return pq
}

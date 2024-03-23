// priority_queue_test.go

package priority_queue

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	frequencies := map[byte]int{
		'a': 3,
		'b': 2,
		'c': 1,
	}
	pq := NewPriorityQueue(frequencies)

	// Ensure items are popped in priority order
	expectedPriority := []int{1, 2, 3}
	for _, priority := range expectedPriority {
		item := heap.Pop(&pq).(*Node)
		if item.Priority != priority {
			t.Errorf("Expected priority %d, got %d", priority, item.Priority)
		}
	}
}

func TestPriorityQueueLen(t *testing.T) {
	frequencies := map[byte]int{
		'a': 3,
		'b': 2,
		'c': 1,
	}
	pq := NewPriorityQueue(frequencies)

	// Ensure Len() returns the correct length
	if pq.Len() != len(frequencies) {
		t.Errorf("Expected length %d, got %d", len(frequencies), pq.Len())
	}
}

package collections

import (
	"testing"

	"github.com/danduggan98/BrowserCoin-Go/core/src/collections"
)

func TestQueue(t *testing.T) {

	// Create new queue
	queue := collections.NewQueue()

	// Check Size() method
	if (queue.Size() != 0) {
		t.Error("Initial queue size should be zero")
	}

	// Add some items, then check size
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	
	if (queue.Size() != 3) {
		t.Errorf("Queue should have %v items, not %v", 3, queue.Size())
	}

	// Pop items, then check size and integrity
	first := queue.Pop()
	second := queue.Pop()

	if (first != 1) {
		t.Errorf("First item in queue should be %v, not %v", 1, first)
	}

	if (second != 2) {
		t.Errorf("Second item in queue should be %v, not %v", 2, second)
	}

	if (queue.Size() != 1) {
		t.Errorf("Queue should have %v items, not %v", 1, queue.Size())
	}

	// Clear queue, then check size
	queue.Clear()

	if (queue.Size() != 0) {
		t.Error("After clearing, queue size should be zero")
	}

	// Popping an empty queue should return nil
	nil_item := queue.Pop()

	if (nil_item != nil) {
		t.Error("After clearing, pop should return nil")
	}
}
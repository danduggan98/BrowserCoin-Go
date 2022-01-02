package collections

type Queue struct {
	data []interface{}
}

// Queue constructor
func NewQueue() *Queue {
	return &Queue{
		data: make([]interface{}, 0),
	}
}

// Remove all items from the queue
func (Q *Queue) Clear() {
	Q.data = make([]interface{}, 0)
}

// Returns the number of items in the queue
func (Q *Queue) Size() int {
	return len(Q.data)
}

// Adds an item to the queue
func (Q *Queue) Push(item interface{}) {
	Q.data = append(Q.data, item)
}

// Removes and returns the oldest item from the queue
func (Q *Queue) Pop() interface{} {
	
	if Q.Size() == 0  {
		return nil
	}

	item := Q.data[0]
	Q.data = Q.data[1:]

	return item
}
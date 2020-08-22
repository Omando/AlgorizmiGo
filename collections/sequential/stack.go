package sequential

import "errors"

/* Stack implementation using an array as the underlying data structure */

// Model the value stored in the Stack as a Item struct. The Item struct has
// a value and may additionally have any requires meta data such date/time
// key, id, etc.
type Item struct {
	Value int
}

type Stack struct {
	data  []Item // underlying data
	count int    // count of items already added to data array
}

// Note that we do not specify a receiver
func CreateStack(capacity int) Stack {
	return Stack{data: make([]Item, capacity), count: 0}
}

func (s *Stack) IsEmpty() bool {
	return s == nil || s.count == 0
}

func (s *Stack) Push(item Item) {
	// If we're at capacity, resize array
	if len(s.data) == s.count {
		// Allocate a new array twice the size of the existing one.
		// TODO: Throttle the increase so as to decay exponentially as the size increases
		// TODO: For example, increase by 100%, then 50%, then 20%, etc.. and finally 10%
		var newData []Item = make([]Item, len(s.data)*2)

		// Copy data to new array then make the new array the Stack's underlying data
		copy(newData, s.data)
		s.data = newData
	}

	// We have capacity. Add item at the next available array cell
	s.data[s.count] = item
	s.count++
}

func (s *Stack) Pop() (Item, error) {
	// Error if Stack is empty
	if s.IsEmpty() {
		return Item{}, errors.New("Stack is empty")
	}

	// Return the item most recently added
	var data Item = s.data[s.count-1]
	s.count--
	return data, nil
}

func (s *Stack) Size() int {
	return s.count
}

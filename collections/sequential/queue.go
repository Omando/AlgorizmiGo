package sequential

import "errors"

/* Queue implementation using a linked list as the underlying data structure */

type node struct {
	data interface{}
	next *node
}

type Queue struct {
	head *node
	tail *node
	size int
}

func CreateQueue() Queue {
	// An empty Queue is one where head points to tail
	return Queue{head: nil, tail: nil, size: 0}
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Enqueue(data interface{}) {
	// Create a new node and append at the end
	var newNode node = node{data: data, next: nil}

	// If the Queue is empty, head points to this new entry
	if q.IsEmpty() {
		q.head = &newNode
		q.tail = &newNode
	} else {
		// Assign the new tail
		q.tail.next = &newNode
		q.tail = &newNode
	}

	// Increment size
	q.size++
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}

	// Head points to the first element. Get its data
	var data = q.head.data

	// Advance head to the next element and decrement size
	q.head = q.head.next
	q.size--

	return data, nil
}

func (q *Queue) Size() int {
	return q.size
}

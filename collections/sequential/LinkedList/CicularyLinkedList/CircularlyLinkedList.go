package CircularlyLinkedList

import "errors"

type CircularlyNode struct {
	value interface{}
	next  *CircularlyNode
}

/* A circularly linked list is essentially a singularly linked list with the following
modifications:
	'next' points back to the head of the list rather than null.
	'head' is not maintained separately since head = tail.next.
	An additional rotate() method which moves the first element to the end of the list.
*/
type CircularlyLinkedList struct {
	tail *CircularlyNode
	size int
}

func New() *CircularlyLinkedList {
	return &CircularlyLinkedList{
		tail: nil,
		size: 0,
	}
}

func (list *CircularlyLinkedList) First() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("list is empty")
	}
	return list.tail.next.value, nil
}

func (list *CircularlyLinkedList) Last() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("list is empty")
	}

	return list.tail.value, nil
}

func (list *CircularlyLinkedList) Length() int {
	return list.size
}

func (list *CircularlyLinkedList) Prepend(value interface{}) {
	newNode := &CircularlyNode{value: value, next: nil}

	if list.size == 0 {
		// list is empty: tail points back to itself.
		list.tail = newNode
		list.tail.next = list.tail
	} else {
		// List not empty. Tail points to the new node
		newNode.next = list.tail.next
		list.tail.next = newNode
	}

	list.size++
}

func (list *CircularlyLinkedList) Append(value interface{}) {
	list.Prepend(value)
	rotate(list)
}

func (list *CircularlyLinkedList) Remove(value interface{}) (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("list is empty")
	}

	// Iterate until we find a node with the given given value
	var previous *CircularlyNode
	tailVisited := false
	for node := list.tail.next; !tailVisited; node = node.next {
		if node.value == value {
			if list.tail.next == node { // Removing head
				list.tail.next = node.next
			} else if list.tail == node { // Removing tail
				previous.next = node.next
				list.tail = previous
			} else { // Everything in between
				previous.next = node.next
			}

			// Update size and reset head/tail if list is now empty
			list.size--
			if list.size == 0 {
				list.tail = nil
			}
			return value, nil // We're done
		}

		// Exit if we have performed a full cycle and value was not found
		if list.tail == node {
			tailVisited = true
		}
		previous = node
	}

	return nil, errors.New("value not found")

}

func (list *CircularlyLinkedList) Contains(value interface{}) bool {
	if list.size == 0 {
		return false
	}

	// Check head and tail first
	firstNode := list.tail.next
	lastNode := list.tail
	if firstNode.value == value || lastNode.value == value {
		return true
	}

	// Iterate through the rest of the list
	for node := firstNode.next; node != lastNode; node = node.next {
		if node.value == value {
			return true
		}
	}

	return false
}

func rotate(list *CircularlyLinkedList) {
	if list.size == 0 {
		return
	}

	list.tail = list.tail.next
}

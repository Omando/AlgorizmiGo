package singlyLinkedList

import (
	"errors"
)

// LinkedListNode node for a singlye linked list
type LinkedListNode struct {
	value interface{}     // Any type. No run-time type-safety. Need generics!
	next  *LinkedListNode // Must use pointer for recursive types. Otherwise size of struct is unknown to compiler
}

// SinglyLinkedList singly linked list
type SinglyLinkedList struct {
	head *LinkedListNode
	tail *LinkedListNode
	size int
}

// Factory function
func New() *SinglyLinkedList {
	// Recall that new returns a pointer to a newly allocated zero value of the given type
	// Can also use:
	//	return &singlyLinkedList{head: nil, tail: nil, size: 0}
	//	return (&singlyLinkedList{}).init()
	return new(SinglyLinkedList).init() // same as:
}

// Receiver is a pointer which receives a copy of the address. If receiver was value,
// a cop of the value is passed and changes inside this function are not visible
// outside the function
func (list *SinglyLinkedList) init() *SinglyLinkedList {
	list.head = nil
	list.tail = nil
	list.size = 0
	return list
}

// First returns the first element of list l or nil if the list is empty.
func (list *SinglyLinkedList) First() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("list is empty")
	}
	return list.head.value, nil
}

// Last returns the last element of list l or nil if the list is empty.
func (list *SinglyLinkedList) Last() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("list is empty")
	}
	return list.tail.value, nil
}

// Length returns count of elements in the list
func (list *SinglyLinkedList) Length() int {
	return list.size
}

// Prepend adds the given value to the beginning of the list
func (list *SinglyLinkedList) Prepend(value interface{}) {
	node := &LinkedListNode{value: value, next: nil}
	// If the list is empty, head and tail both point to the new node
	if list.size == 0 {
		list.head = node
		list.tail = node
	} else {
		node.next = list.head // new node points to the current head
		list.head = node      // new node is the new head
	}

	// Update size
	list.size++
}

// Append adds the given value to the end of the list
func (list *SinglyLinkedList) Append(value interface{}) {
	node := &LinkedListNode{value: value, next: nil}
	// If the list is empty, head and tail both point to the new node
	if list.size == 0 {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}
	// Update size
	list.size++
}

// Remove deletes the given node if found, and returns value of removed node.
// Panics if value was not found
func (list *SinglyLinkedList) Remove(value interface{}) (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("list is empty")
	}

	// Iterate until we find a node with the given given value
	var previous *LinkedListNode
	for node := list.head; node != nil; node = node.next {
		if node.value == value {
			if list.head == node { // Removing head
				list.head = list.head.next
			} else if list.tail == node { // Removing tail
				previous.next = nil
				list.tail = previous
			} else { // Everything in between
				previous.next = node.next
			}

			// Update size and reset head/tail if list is now empty
			list.size--
			if list.size == 0 {
				list.head = nil
				list.tail = nil
			}
			return value, nil
		}
		previous = node
	}

	return nil, errors.New("value not found")
}

// Contains returns true if given value exists, otherwise returns false
func (list *SinglyLinkedList) Contains(value interface{}) bool {
	if list.size == 0 {
		return false
	}

	// Check head and tail first
	if list.head.value == value || list.tail.value == value {
		return true
	}

	// Iterate through the rest of the list
	for node := list.head.next; node != list.tail; node = node.next {
		if node.value == value {
			return true
		}
	}

	return false
}

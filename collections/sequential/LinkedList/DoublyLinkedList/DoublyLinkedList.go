package DoublyLinkedList

import (
	"errors"
)

type DoublyNode struct {
	value    interface{} // Any type. No run-time type-safety. Need generics!
	next     *DoublyNode // Must use pointer for recursive types. Otherwise size of struct is unknown to compiler
	previous *DoublyNode
}

// DoublyLinkedList doubly linked list with head and tail sentinels
type DoublyLinkedList struct {
	headSentinel *DoublyNode
	tailSentinel *DoublyNode
	size         int
}

// New Initializes an empty doubly linked list
func New() *DoublyLinkedList {
	list := &DoublyLinkedList{
		headSentinel: new(DoublyNode),
		tailSentinel: new(DoublyNode),
		size:         0,
	}

	// Empty linked list: head and tail sentinels point at each other
	list.headSentinel.next = list.tailSentinel
	list.tailSentinel.previous = list.headSentinel

	return list
}

// First returns the first element of list l or nil if the list is empty.
func (list *DoublyLinkedList) First() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("List is empty")
	}

	// Note use of headSentinel.next
	return list.headSentinel.next.value, nil
}

// Last returns the last element of list l or nil if the list is empty.
func (list *DoublyLinkedList) Last() (interface{}, error) {
	if list.size == 0 {
		return nil, errors.New("List is empty")
	}

	// Note use of tailSentinel.previous
	return list.tailSentinel.previous.value, nil
}

// Length returns count of elements in the list
func (list *DoublyLinkedList) Length() int {
	return list.size
}

// Prepend adds the given value to the beginning of the list
func (list *DoublyLinkedList) Prepend(element interface{}) {
	addBetween(list, list.headSentinel, list.headSentinel.next, element)
}

// Append adds the given value to the end of the list
func (list *DoublyLinkedList) Append(element interface{}) {
	addBetween(list, list.tailSentinel.previous, list.tailSentinel, element)
}

// Remove deletes the given node if found, and returns value of removed node.
// Panics if value was not found
// This function is best understood by drawing nodes and noting connections:
// 1) Draw and connect three double linked nodes (next, node to delete, previous) before
// deleting the given value
// 2) Draw and connect two double linked nodes after deleting the node for the given value
func (list *DoublyLinkedList) Remove(element interface{}) (interface{}, error) {
	// Locate the node first
	nodeToDelete, err := findNode(list, element)
	if err != nil {
		return nil, err
	}

	previousNode := nodeToDelete.previous
	nextNode := nodeToDelete.next
	previousNode.next = nextNode
	nextNode.previous = previousNode

	// Decrement size
	list.size--

	// Return value of the element that was removed
	return nodeToDelete.value, nil
}

// Contains returns true if given value exists, otherwise returns false
func (list *DoublyLinkedList) Contains(value interface{}) bool {
	if list.size == 0 {
		return false
	}

	// Check head and tail first
	firstNode := list.headSentinel.next
	lastNode := list.tailSentinel.previous
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

// This function is best understood by drawing nodes and noting connections:
// 1) Draw and connect two double linked nodes before adding a node in between with the given value
// 1) Draw and connect three double linked nodes after adding the new node
func addBetween(list *DoublyLinkedList, leftNode *DoublyNode, rightNode *DoublyNode, value interface{}) {
	newNode := &DoublyNode{
		value:    value,
		next:     rightNode,
		previous: leftNode,
	}

	leftNode.next = newNode
	rightNode.previous = newNode

	// finally, increment size
	list.size++
}

func findNode(list *DoublyLinkedList, value interface{}) (*DoublyNode, error) {
	for node := list.headSentinel.next; node != nil; node = node.next {
		if node.value == value {
			return node, nil
		}
	}

	return nil, errors.New("Value not found")
}

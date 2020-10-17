package linkedLists

type Node struct {
	Value int
	Next  *Node // Must use pointer for recursive types. Otherwise size of struct is unknown to compiler
}

// Recall that the zero value of a struct is a struct with all fields set to
// their own zero values. This means a struct cannot be compared to nil.
// To allow comparison with nil, we declare head and tail struct as pointers
type SimpleSinglyLinkedList struct {
	head *Node
	tail *Node
}

// Recall: A non-pointer receiver does not show changes to the receiver outside the function
func (ll *SimpleSinglyLinkedList) Head() *Node {
	return ll.head
}

// Recall: A non-pointer receiver does not show changes to the receiver outside the function
func (ll *SimpleSinglyLinkedList) Append(data int) {
	newNode := &Node{data, nil}
	// Head and tail point to the same node if the list is empty
	if ll.head == nil {
		ll.head = newNode
		ll.tail = ll.head
		return
	}

	// List is already populated. Append a new node to the end
	ll.tail.Next = newNode
	ll.tail = newNode
}

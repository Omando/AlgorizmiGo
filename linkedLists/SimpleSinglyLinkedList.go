package linkedLists

type Node struct {
	value int
	next *Node	// Must use pointer for recursive types. Otherwise size of struct is unknown to compiler
}

// Recall that the zero value of a struct is a struct with all fields set to
// their own zero values. This means a struct cannot be compared to nil
// To allow comparison with nil, we declare head and tail struct as pointers
type SimpleSinglyLinkedList struct {
	head *Node
	tail *Node
}

func (linkedList SimpleSinglyLinkedList) append(data int) {
	newNode := &Node{value: data, next: nil}
	// Head and tail point to the same node if the list is empty
	if linkedList.head == nil {
		linkedList.head = newNode
		linkedList.tail = linkedList.head
	}

	// List is already populated. Append a new node to the end
	linkedList.tail.next = newNode
	linkedList.tail = newNode
}

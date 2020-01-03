package recursion

type Node struct {
	Value int
	Next *Node
}
func SwapLinkedListNodes(head *Node) *Node {
	// Must append a dummy
	var dummy *Node = &Node{Value: 0, Next: head}
	doSwap(dummy, head, head.Next)
	return dummy.Next
}

func doSwap(leftleft *Node, left *Node, right *Node) {
	// Exit condition
	if left == nil || right == nil {
		return
	}

	// Perform swap
	leftleft.Next = right;
	left.Next = right.Next;
	right.Next = left;

	// Next pair
	if left.Next != nil && left.Next.Next != nil {
		doSwap(left, left.Next, left.Next.Next)
	}
}

package recursion

import "AlgorizmiGo/linkedLists"

func SwapLinkedListNodes(head *linkedLists.Node) *linkedLists.Node {
	// Must append a dummy
	var dummy *linkedLists.Node = &linkedLists.Node{Value: 0, Next: head}
	doSwap(dummy, head, head.Next)
	return dummy.Next
}

func doSwap(leftleft *linkedLists.Node, left *linkedLists.Node, right *linkedLists.Node) {
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

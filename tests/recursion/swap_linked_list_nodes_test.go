package recursion

import (
	"AlgorizmiGo/linkedLists"
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwappingLinkedListNodes(t *testing.T) {
	tests := []struct {
		head                      *linkedLists.Node
		expectedSwappedLinkedList *linkedLists.Node
	}{
		{head: getLinkedList1(), expectedSwappedLinkedList: getSwappedLinkedList1()},
	}

	for _, test := range tests {
		actualSwappedLinkedList := recursion.SwapLinkedListNodes(test.head)

		currentExpectedNode := test.expectedSwappedLinkedList
		currentActualNode := actualSwappedLinkedList
		for currentExpectedNode.Next != nil && currentActualNode != nil {
			assert.EqualValues(t, currentExpectedNode.Value, currentActualNode.Value)
			currentExpectedNode = currentExpectedNode.Next
			currentActualNode = currentActualNode.Next
		}
	}
}

func getLinkedList1() *linkedLists.Node {
	ll := linkedLists.SimpleSinglyLinkedList{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	return ll.Head()
}

func getSwappedLinkedList1() *linkedLists.Node {
	ll := linkedLists.SimpleSinglyLinkedList{}
	ll.Append(2)
	ll.Append(1)
	ll.Append(4)
	ll.Append(3)
	return ll.Head()
}

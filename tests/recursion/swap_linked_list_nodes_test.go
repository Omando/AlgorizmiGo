package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwappingLinkedListNodes(t *testing.T) {
	tests := []struct {
		head                      *recursion.Node
		expectedSwappedLinkedList *recursion.Node
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

func getLinkedList1() *recursion.Node {
	var node4 = recursion.Node{4, nil}
	var node3 = recursion.Node{3, &node4}
	var node2 = recursion.Node{2, &node3}
	var node1 = recursion.Node{1, &node2}

	return &node1
}

func getSwappedLinkedList1() *recursion.Node {
	var node4 = recursion.Node{3, nil}
	var node3 = recursion.Node{4, &node4}
	var node2 = recursion.Node{1, &node3}
	var node1 = recursion.Node{2, &node2}

	return &node1
}

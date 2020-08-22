package sequentialCollections

import (
	"AlgorizmiGo/collections/sequential"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStackIsEmpty(t *testing.T) {
	// Arrange
	var stack = sequential.CreateStack(10)

	// Act
	assert.Equal(t, stack.IsEmpty(), true)
}

func TestShouldPushItems(t *testing.T) {
	tests := []struct {
		name  string
		data  []int
		count int
	}{
		{name: "test1", data: []int{}, count: 0},
		{name: "test2", data: []int{1, 2, 3, 4}, count: 4},
		{name: "test3", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, count: 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var stack sequential.Stack = sequential.CreateStack(4)
			for _, n := range test.data {
				stack.Push(sequential.Item{n})
			}

			if stack.Size() != test.count {
				t.Errorf("stack.Size() :%v, expected:%v", stack.Size(), test.count)
			}
		})
	}
}

func TestShouldPopItems(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{name: "test1", input: nil, expectedOutput: nil},
		{name: "test2", input: []int{1, 2, 3}, expectedOutput: []int{3, 2, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var stack sequential.Stack = sequential.CreateStack(5)
			for _, n := range test.input {
				stack.Push(sequential.Item{n})
			}

			// Collect Pop outputs
			var actualOutput []int // slice!
			for !stack.IsEmpty() { // conditional while
				data, err := stack.Pop()
				if err == nil {
					actualOutput = append(actualOutput, data.Value)
				}
			}

			// Compare
			assert.Equal(t, test.expectedOutput, actualOutput)
		})
	}
}

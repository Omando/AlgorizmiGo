package sequentialCollections

import (
	"AlgorizmiGo/collections/sequential"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueueIsEmpty(t *testing.T) {
	// Arrange
	var queue = sequential.CreateQueue()

	// Act
	assert.Equal(t, queue.IsEmpty(), true)
}

func TestShouldEnqueueItems(t *testing.T) {
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
			var queue sequential.Queue = sequential.CreateQueue()
			for _, n := range test.data {
				queue.Enqueue(n)
			}

			if queue.Size() != test.count {
				t.Errorf("queue.Size() :%v, expected:%v", queue.Size(), test.count)
			}
		})
	}
}

func TestShouldDequeueItems(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{name: "test1", input: nil, expectedOutput: nil},
		{name: "test2", input: []int{1, 2, 3}, expectedOutput: []int{1, 2, 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var queue sequential.Queue = sequential.CreateQueue()
			for _, n := range test.input {
				queue.Enqueue(n)
			}

			// Collect Pop outputs
			var actualOutput []int // slice!
			for !queue.IsEmpty() { // conditional while
				data, err := queue.Dequeue()
				if err == nil {
					actualOutput = append(actualOutput, data)
				}
			}

			// Compare
			assert.Equal(t, test.expectedOutput, actualOutput)
		})
	}
}

package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetSum(t *testing.T) {
	tests := []struct {
		input          []int
		targetSum      int
		expectedOutput bool
	}{
		{[]int{2, 3, 4}, 4, true},
		{[]int{2, 1, 3, 4}, 4, true},
		{[]int{2, 1, 3, 4}, 2, true},
		{[]int{2, 1, 3, 4}, 7, true},
		{[]int{2, 1, 3, 4}, 9, true},
		{[]int{2, 1, 3, 4}, 10, true},
		{[]int{2, 1, 3, 4}, 8, true},
		{[]int{2, 1, 3, 4}, 100, false},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 20, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 12, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 50, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 51, false},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 49, false},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 48, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 26, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 27, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 29, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 31, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 33, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 33, true},
		{[]int{2, 8, 4, 6, 15, 10, 5}, 47, false},
	}

	for _, test := range tests {
		actualOutput := recursion.SubsetSum(test.input, test.targetSum)
		assert.Equal(t, test.expectedOutput, actualOutput)
	}
}

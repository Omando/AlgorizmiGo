package dynamicProgamming

import (
	"AlgorizmiGo/dynamicProgramming"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_calcualte_optimal_costs_and_roots(t *testing.T) {
	tests := []struct {
		name          string
		keys          []byte
		frequencies   []int
		expectedCosts [][]int
		expectedRoots [][]int
	}{
		{
			name:        "5 Keys",
			keys:        []byte{'A', 'B', 'C', 'D', 'E'},
			frequencies: []int{25, 20, 5, 20, 30},
			expectedCosts: [][]int{
				{25,65,80,125,210},
				{0, 20, 30, 75, 135},
				{0, 0, 5, 30, 85},
				{0, 0, 0, 20, 70},
				{0, 0, 0, 0, 30},
			},
			expectedRoots: [][]int{
				{0, 0, 0, 1, 1},
				{0, 1, 1, 1, 3},
				{0, 0, 2, 3, 4},
				{0, 0, 0, 3, 4},
				{0, 0, 0, 0, 4},
			},
		},
	}
	for _, test := range tests {
		costs, roots, err := dynamicProgramming.FindOptimalBST(test.keys, test.frequencies)
		assert.Nil(t, err)
		assert.EqualValues(t, costs, test.expectedCosts)
		assert.EqualValues(t, roots, test.expectedRoots)
	}
}

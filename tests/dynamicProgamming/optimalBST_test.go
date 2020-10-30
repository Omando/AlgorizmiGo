package dynamicProgamming

import (
	"AlgorizmiGo/dynamicProgramming"
	"github.com/stretchr/testify/assert"
	"testing"
)

func should_calcualte_optimal_costs_and_roots(t *testing.T) {
	tests := []struct {
		name string
		keys []byte
		frequencies []int
		expectedCosts [][]int
		expectedRoots [][]int
	}{
		{
			name: 		   "5 Keys",
			keys:          []byte{'A','B','C','D','E'},
			frequencies:   []int{25,20,5,20,30},
			expectedCosts: [][]int {
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,0,0,0,0},
			},
			expectedRoots: [][]int {
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,0,0,0,0},
			},
		},
	}
	for _, test := range tests {
		costs, roots, err := dynamicProgramming.FindOptimalBST(test.keys, test.frequencies)
		assert.NotNil(t,err)
		assert.EqualValues(t, costs, test.expectedCosts)
		assert.EqualValues(t, roots, test.expectedRoots)
	}
}

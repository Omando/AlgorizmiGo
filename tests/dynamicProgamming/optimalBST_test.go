package dynamicProgamming

import (
	"AlgorizmiGo/dynamicProgramming"
	"fmt"
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
		t.Run(test.name, func(t *testing.T) {
			costs, roots := dynamicProgramming.FindOptimalBST(test.keys, test.frequencies)
			fmt.Println(costs, roots);
		})
	}
}

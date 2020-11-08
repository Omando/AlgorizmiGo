package dynamicProgamming

import (
	"AlgorizmiGo/dynamicProgramming"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_calcualte_subset_for_given_sum(t *testing.T) {
	tests := []struct {
		input          []int
		sum            int
		expectedSubset []int
	}{
		{
			input:          []int{1, 4, 2, 3},
			sum:            5,
			expectedSubset: []int{4, 1},
		},
		{
			input:          []int{2, 4, 6, 7, 8, 4, 10, 5, 24, 12},
			sum:            35,
			expectedSubset: []int{10, 8, 7, 6, 4},
		},
	}

	for _, test := range tests {
		actualSubset, err := dynamicProgramming.SubsetSum(test.input, test.sum)
		assert.Nil(t, err)
		assert.EqualValues(t, test.expectedSubset, actualSubset)
	}
}

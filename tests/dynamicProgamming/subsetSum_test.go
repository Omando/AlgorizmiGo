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
	}

	for _, test := range tests {
		actualSubset, err := dynamicProgramming.SubsetSum(test.input, test.sum)
		assert.Nil(t, err)
		assert.EqualValues(t, test.expectedSubset, actualSubset)
	}
}

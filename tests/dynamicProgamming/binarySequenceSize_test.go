package dynamicProgamming

import (
	"AlgorizmiGo/dynamicProgramming"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_calculate_binary_sequence_size(t *testing.T) {
	tests := []struct {
		n            int
		expectedSize int
	}{
		{n: 1, expectedSize: 2},
		{n: 2, expectedSize: 4},
		{n: 3, expectedSize: 8},
		{n: 5, expectedSize: 32},
	}

	for _, test := range tests {
		actualSize, err := dynamicProgramming.CalculateBinarySequenceSize(test.n)
		assert.Nil(t, err)
		assert.Equal(t, test.expectedSize, actualSize)
	}
}

func Test_should_calculate_binary_sequence_no_consecutive_zeros_size(t *testing.T) {
	tests := []struct {
		n            int
		expectedSize int
	}{
		{n: 1, expectedSize: 2},
		{n: 2, expectedSize: 3},
		{n: 3, expectedSize: 5},
		{n: 4, expectedSize: 8},
		{n: 5, expectedSize: 13},
	}

	for _, test := range tests {
		actualSize, err := dynamicProgramming.CalculateNoConsecutiveZerosBinarySequenceSize(test.n)
		assert.Nil(t, err)
		assert.Equal(t, test.expectedSize, actualSize)
	}
}

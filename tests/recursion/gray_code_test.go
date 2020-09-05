package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrayCode(t *testing.T) {
	tests := []struct {
		n              int
		expectedOutput []string
	}{
		{n: 2, expectedOutput: []string{"00", "01", "11", "10"}},
		{n: 3, expectedOutput: []string{"000", "001", "011", "010", "110", "111", "101", "100"}},
		{n: 4, expectedOutput: []string{"0000", "0001", "0011", "0010", "0110", "0111", "0101", "0100",
			"1100", "1101", "1111", "1110", "1010", "1011", "1001", "1000"}},
	}

	for _, test := range tests {
		actualOutput := recursion.GrayCode(test.n)
		assert.Equal(t, test.expectedOutput, actualOutput)
	}
}

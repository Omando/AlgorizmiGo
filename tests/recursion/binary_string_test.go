package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryStrings(t *testing.T) {
	tests := []struct {
		number         int
		expectedOutput string
	}{
		{number: 0, expectedOutput: "0"},
		{number: 1, expectedOutput: "1"},
		{number: 2, expectedOutput: "10"},
		{number: 10, expectedOutput: "1010"},
		{number: 15, expectedOutput: "1111"},
	}

	for _, test := range tests {
		actualOutput, err := recursion.BinaryString(test.number)
		assert.Nil(t, err)
		assert.Equal(t, test.expectedOutput, actualOutput)
	}
}

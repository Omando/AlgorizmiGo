package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringPermutations(t *testing.T) {
	tests := []struct {
		input                string
		expectedPermutations []string
	}{
		{input: "a", expectedPermutations: []string{"a"}},
		{input: "ab", expectedPermutations: []string{"ab", "ba"}},
		{input: "abc", expectedPermutations: []string{"abc", "acb", "bca", "bac", "cab", "cba"}},
		{input: "يزن", expectedPermutations: []string{"يزن", "ينز", "زني", "زين", "نيز", "نزي"}},
	}

	for _, test := range tests {
		actualPermutations := recursion.StringPermutations(test.input)
		assert.ElementsMatch(t, test.expectedPermutations, actualPermutations)
	}
}

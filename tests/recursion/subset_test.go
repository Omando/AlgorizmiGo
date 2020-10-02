package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubset(t *testing.T) {
	tests := []struct {
		input	string
		expectedSubsets []string
	}{
		{"ab", []string{"a", "b", "ab"}},
		{"abc", []string{"a","b", "c", "ab", "ac", "bc", "abc"}},
		{"يس", []string{"س","ي", "يس"}},
	}
	for _, test := range tests {
		actualSubset := recursion.Subset(test.input)
		assert.ElementsMatch(t, test.expectedSubsets, actualSubset)
	}
}

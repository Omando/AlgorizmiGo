package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCharacterCount(t *testing.T) {
	inputs := []struct {
		input         string
		char          byte
		expectedCount int
	}{
		{"", 'a', 0},
		{"Z", 'a', 0},
		{"a", 'a', 1},
		{"hello world", 'o', 2},
		{"ababab acacac", 'a', 6},
		{"xxx yyy xxx", 'x', 6},
	}

	for _, input := range inputs {
		actualCount := recursion.GetCharacterCount(input.input, input.char)
		assert.Equal(t, input.expectedCount, actualCount)
	}
}

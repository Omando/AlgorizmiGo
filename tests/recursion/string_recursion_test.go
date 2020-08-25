package recursion

import (
	"AlgorizmiGo/recursion"
	"fmt"
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

func TestRemoveVowels(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{"bcdf", "bcdf"},
		{"abcdefg", "bcdfg"},
		{"iiisssooo", "sss"},
		{"aeiou", ""},
	}

	for _, test := range tests {
		actualOutput := recursion.RemoveVowels(test.input)
		assert.Equal(t, test.output, actualOutput)
	}
}

func TestReverseString(t *testing.T) {
	const myString = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Printf("%q\n", myString)

	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{"bcdf", "fdcb"},
		{"abcdefg", "gfedcba"},
		{"qwerty", "ytrewq"},
		{"yazan diranieh", "heinarid nazay"},
	}

	for _, test := range tests {
		actualOutput := recursion.ReverseString(test.input)
		assert.Equal(t, test.output, actualOutput)
	}
}

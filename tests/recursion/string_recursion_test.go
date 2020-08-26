package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"sort"
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

func TestAllStringCombinations(t *testing.T) {
	tests := []struct {
		input string
		expectedCombinations []string
	} {
		{ input: "", expectedCombinations: []string{},},
		{ input: "a", expectedCombinations: []string{"a"}},
		{ input:"ab", expectedCombinations: []string{"b", "a", "ab"}},
		{ input:"ab⌘", expectedCombinations: []string{"b", "a⌘", "a", "ab", "ab⌘", "b⌘", "⌘"}},
	}

	for _, test := range tests {
		actualCombinations := recursion.AllStringCombinations(test.input)
		sort.Strings(actualCombinations)
		sort.Strings(test.expectedCombinations)
		assert.Equal(t, test.expectedCombinations, actualCombinations)
	}
}


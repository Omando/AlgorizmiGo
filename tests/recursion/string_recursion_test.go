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
		input                string
		expectedCombinations []string
	}{
		{input: "", expectedCombinations: []string{}},
		{input: "a", expectedCombinations: []string{"a"}},
		{input: "ab", expectedCombinations: []string{"b", "a", "ab"}},
		{input: "ab⌘", expectedCombinations: []string{"b", "a⌘", "a", "ab", "ab⌘", "b⌘", "⌘"}},
	}

	for _, test := range tests {
		actualCombinations := recursion.AllStringCombinations(test.input)
		sort.Strings(actualCombinations)
		sort.Strings(test.expectedCombinations)
		assert.Equal(t, test.expectedCombinations, actualCombinations)
	}
}

func TestStringInterleaves(t *testing.T) {
	tests := []struct {
		s1                  string
		s2                  string
		expectedInterleaves []string
	}{
		{s1: "A", s2: "C", expectedInterleaves: []string{"AC", "CA"}},
		{s1: "AB", s2: "C", expectedInterleaves: []string{"ABC", "ACB", "CAB"}},
		{s1: "AB", s2: "CD", expectedInterleaves: []string{"ABCD", "ACBD", "ACDB", "CDAB", "CABD", "CADB"}},
	}

	for _, test := range tests {
		actualInterleaves := recursion.InterleaveStrings(test.s1, test.s2)
		sort.Strings(actualInterleaves)
		sort.Strings(test.expectedInterleaves)
		assert.Equal(t, test.expectedInterleaves, actualInterleaves)
	}
}

func TestCombinationsFromKeyPad(t *testing.T) {
	tests := []struct {
		numbers              []int
		expectedCombinations []string
	}{
		{numbers: []int{1}, expectedCombinations: []string{""}},
		{numbers: []int{1, 2}, expectedCombinations: []string{"A", "B", "C"}},
		{numbers: []int{2, 3}, expectedCombinations: []string{"AD", "AE", "AF", "BD", "BE", "BF", "CD", "CE", "CF"}},
		{numbers: []int{2, 3, 4}, expectedCombinations: []string{
			"ADG", "ADH", "ADI", "AEG", "AEH", "AEI", "AFG", "AFH", "AFI",
			"BDG", "BDH", "BDI", "BEG", "BEH", "BEI", "BFG", "BFH", "BFI",
			"CDG", "CDH", "CDI", "CEG", "CEH", "CEI", "CFG", "CFH", "CFI"}},
	}

	for _, test := range tests {
		actualCombinations := recursion.CombinationsFromDialKeypad(test.numbers)
		assert.Equal(t, test.expectedCombinations, actualCombinations)
	}
}

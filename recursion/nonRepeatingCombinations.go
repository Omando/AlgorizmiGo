package recursion

import (
	"strings"
)

var combinations []string

func GetNonRepeatingCombinations(universe []string) []string {
	var currentSequence strings.Builder
	// Builder must not be copied (i.e, passed by value). Otherwise you get
	// the following error when getCombinations is called recursively:
	// illegal use of non-zero Builder copied by value
	getCombinations(universe, &currentSequence)
	return combinations
}

func getCombinations(universe []string, currentSequence *strings.Builder) {
	// Exit condition
	if len(universe) == 0 {
		combinations = append(combinations, currentSequence.String())
		return
	}

	// Iterate over each character in the universe
	for index, value := range universe {
		// Add this character
		currentSequence.WriteString(value)

		// For the next level, the universe of characters should exclude the current
		// character, i.e., remove character at index i from universe
		universe = append(universe[:index], universe[index+1:]...) // Remove element at the given index
		getCombinations(universe, currentSequence)

		// Re-add the excluded character back to the universe because the next level
		// has now been processed
		universe = append(universe, "")            // zero value for string is ""
		copy(universe[index+1:], universe[index:]) // copy(destination, source). Shift elements from index to index+!
		universe[index] = value                    // Copy a value into the vacated element at the given index

		// Remove the last character added because the next level has now been processed
		var currentSequenceAsArray []int32 = []int32(currentSequence.String())
		currentSequence.Reset()
		currentSequence.WriteString(string(currentSequenceAsArray[:len(currentSequenceAsArray)-1]))
	}
}

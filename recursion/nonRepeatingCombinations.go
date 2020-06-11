package recursion

import (
	"strings"
)

func GetNonRepeatingCombinations(universe []string) []string {
	// TODO
	return nil
}

func getCombinations(universe []string, combinations []string, currentSequence strings.Builder) {
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
		universe = append(universe[:index], universe[index+1:]...)
		getCombinations(universe, combinations, currentSequence)

		// Re-add the excluded character back to the universe because the next level
		// has now been processed
		// TODO:  INSERT VALUE AT INDEX in universe

		// Remove the last character added because the next level has now been processed
		// (look at NonRepeatingCombination diagram)
		// TODO:
	}
}

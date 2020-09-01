package recursion

import (
	"strings"
)

/* Counts how many times a given character appears in a given string */
func GetCharacterCount(s string, c byte) int {
	if len(s) == 0 {
		return 0
	}

	return doGetCharacterCount(s, c, 0, 0)
}

func doGetCharacterCount(s string, c byte, index int, runningCount int) int {
	// Exit condition
	if len(s) == index {
		return runningCount
	}

	// Calculation
	if s[index] == c {
		runningCount++
	}
	return doGetCharacterCount(s, c, index+1, runningCount)
}

/* Returns a new string with all vowels removed. The original string is unchanged
While this code iterates the string from left to right, another approach worth
looking at is iterating over the string from both sides
*/
func RemoveVowels(s string) string {

	// Pass address of string.Builder to avoid copying string builder
	// Otherwise you get the following error:
	// "illegal use of non-zero Builder copied by value"
	return doRemoveVowels(s, 0, &strings.Builder{})
}

func doRemoveVowels(s string, index int, builder *strings.Builder) string {
	// Exit condition
	if len(s) == index {
		return builder.String()
	}

	// Remove vowels if any from both sides by only writing to the string
	currentChar := s[index]
	if currentChar != 'a' &&
		currentChar != 'e' &&
		currentChar != 'i' &&
		currentChar != 'o' &&
		currentChar != 'u' {
		builder.WriteByte(currentChar)
	}

	return doRemoveVowels(s, index+1, builder)
}

/*  Return a reversed string */
func ReverseString(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Convert string to a character array
	chars := []byte(s)
	doReverseString(chars, 0)

	// Convert reversed character array back to string
	return string(chars)
}

// chars slice is not passed as reference because code will work on the underlying
// array and never append
func doReverseString(chars []byte, index int) {
	// Exit condition
	if len(chars)/2 == index {
		return
	}

	// Calculation
	tempChar := chars[index]
	rhsIndex := len(chars) - index - 1
	chars[index] = chars[rhsIndex]
	chars[rhsIndex] = tempChar

	// Recursion
	doReverseString(chars, index+1)
}

/* Given a string, return all possible combinations. For example, given "abc", all possible
character combinations are: "a", "ab", "abc", "ac", "b", "bc", and "c"  */
func AllStringCombinations(input string) []string {

	var runes = []rune(input) // Convert input string to an array of runes
	var combinationSet = make(map[string]bool)

	doAllStringCombinations(runes, combinationSet)

	// Collect keys into a slice
	var combinations []string = make([]string, len(combinationSet))
	index := 0
	for k, _ := range combinationSet {
		combinations[index] = k
		index++
	}
	return combinations
}

func doAllStringCombinations(runes []rune, combinationSet map[string]bool) {
	// Exit condition
	if len(runes) == 0 {
		return
	}

	// Convert rune array to a string and add to set of combinations
	combination := string(runes)
	if !combinationSet[combination] {
		combinationSet[combination] = true
	}

	// Iterate over each character in the string, omitting the current one and
	// repeating the process
	for i := 0; i < len(runes); i++ {
		// Omit current character and get new combinations for the new string
		var tempRune []rune = make([]rune, len(runes))
		copy(tempRune, runes)
		tempRune = append(tempRune[:i], tempRune[i+1:]...)
		doAllStringCombinations(tempRune, combinationSet)
	}
}

/* Given two strings, find all interleaves that can be formed from the characters of
the two strings where order of characters is preserved. For example, given AB and AC,
the possible interleaves are: ABAC, AACB, AABC, ACAB, AABC, AACB */
func InterleaveStrings(s1 string, s2 string) []string {

	var allInterleaves []string

	// allInterleaves slice passed as a reference because it will be appended to
	doInterleaveStrings([]rune(s1), []rune(s2), []rune{}, &allInterleaves)

	return allInterleaves
}

// allInterleaves slice passed as a reference because it will be appended to and
//the internal pointer will change
func doInterleaveStrings(s1 []rune, s2 []rune, result []rune, interleaves *[]string) {

	// Exit condition
	if len(s1) == 0 && len(s2) == 0 {
		*interleaves = append(*interleaves, string(result))
		return
	}

	if len(s1) > 0 {
		newResult := []rune(result)
		newResult = append(newResult, s1[0])
		doInterleaveStrings(s1[1:], s2, newResult, interleaves) // Pass s1 substring starting from index 1
	}

	if len(s2) > 0 {
		newResult := []rune(result)
		newResult = append(newResult, s2[0])
		doInterleaveStrings(s1, s2[1:], newResult, interleaves) // Pass s2 substring starting from index 1
	}
}

/*  Given a sequence of numbers [2..9] from a dial keypad, print all possible combinations
of letters: (2, ABC), (3,DEF), (4, GHI), (5, JKL), (6, MNO), (7, PQRS), (8, TUV), (9,WXYZ)
*/
func combinationsFromDialKeypad(numbers []string) []string {

}

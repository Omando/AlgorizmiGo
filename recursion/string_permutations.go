package recursion

func StringPermutations(input string) []string {
	//  Convert to runes and get count
	inputRunes := []rune(input)
	runeCount := len(inputRunes)

	// Allocate memory for a rune array to hold a full permutation
	var permutation []rune = make([]rune, runeCount)

	// perform actual permutation calculations
	var permutations []string
	calculate(inputRunes, runeCount, 0, permutation, &permutations)
	return permutations
}

func calculate(inputRunes []rune, length int, index int, permutation []rune, permutations *[]string) {
	// Exit condition
	if length == 0 {
		*permutations = append(*permutations, string(permutation))
		return
	}

	// Iterate over all runes for the given length and index
	for i := 0; i < length; i++ {
		permutation[index] = inputRunes[i]

		// For the current character, swap with the last character
		inputRunes[i], inputRunes[length-1] = inputRunes[length-1], inputRunes[i]

		calculate(inputRunes, length-1, index+1, permutation, permutations)

		// For the current character, reset the swap with the last character
		inputRunes[i], inputRunes[length-1] = inputRunes[length-1], inputRunes[i]
	}
}

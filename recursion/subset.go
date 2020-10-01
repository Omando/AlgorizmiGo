package recursion

func Subset(input string) []string {
	// Convert input string to a rune
	inputAsRunes := []rune(input)
	var subsets []string
	var item []rune

	doSubset(inputAsRunes, 0, item, &subsets)

	return subsets
}

func doSubset(input []rune, index int, item []rune, subsets *[]string) {
	// Exit condition
	if len(input) == index {
		return
	}

	newItem := append(item, input[index])
	*subsets = append(*subsets, string(newItem))

	// Branch 1: Include item at current index
	doSubset(input, index+1, newItem, subsets)

	// Branch 1: Exclude item at current index
	doSubset(input, index+1, item, subsets)
}

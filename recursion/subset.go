package recursion

/* Given an array of non-consecutive numbers and a target value, find whether
there exists a subset of those numbers that sums up to the target value*/
func SubsetSum(input []int, targetSum int) bool {
	return doSubSetSum(input, targetSum, 0, 0)
}

func doSubSetSum(input []int, targetSum int, index int, currentSum int) bool {
	if currentSum == targetSum {
		return true
	}

	if currentSum > targetSum || index == len(input) {
		return false
	}

	var hasTargetSum = doSubSetSum(input, targetSum, index+1, currentSum+input[index])
	if hasTargetSum == false {
		hasTargetSum = doSubSetSum(input, targetSum, index+1, currentSum)
	}

	return hasTargetSum
}

func Subset(input string) []string {
	// Convert input string to a rune
	inputAsRunes := []rune(input)
	var subsets []string
	var item []rune

	doSubset(inputAsRunes, 0, item, &subsets )

	return subsets
}

func doSubset(input []rune, index int, item []rune, subsets *[]string) {
	// Exit condition
	if len(input) == index {
		return
	}

	newItem := append(item, input[index])
	*subsets = append(*subsets, string(newItem))

	// Branch 1: Include item at index
	doSubset(input, index+1, newItem, subsets)

	// Branch 1: Exclude item at index
	doSubset(input, index+1, item, subsets)
}
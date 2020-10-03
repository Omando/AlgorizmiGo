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

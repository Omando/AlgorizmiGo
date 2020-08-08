package dynamicProgramming

import "math"

func RodMaximumProfitCut(prices []int, rodLength int) int {
	// No cut is possible if original rod length is zero
	if rodLength == 0 {
		return 0
	}

	var subProblems = make([]int, rodLength+1)
	subProblems[0] = 0 // A rod of length 0 has a price of 0

	// Code is deduced directly from the state-action-destination table on page 26
	// (DSALG_DynamicProgramming). Note the use of state and destination indices that
	// relate directly to state and destination columns and their values
	for state := 1; state <= rodLength; state++ {
		for destination := 0; destination < state; destination++ {
			length := state - destination
			price := prices[length]

			subProblems[state] = int(math.Max(float64(subProblems[state]), float64(price+subProblems[destination])))
		}
	}

	return subProblems[rodLength]
}

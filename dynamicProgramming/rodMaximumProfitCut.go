package dynamicProgramming

import "math"

// Recursion
func RodMaximumProfitCut_Recursive(cutPrices []int, currentLength int) int {
	// No cut is possible if original rod length is zero
	if currentLength == 0 {
		return 0
	}

	// We assume that the current length has the maximum price
	currentMaxPrice := cutPrices[currentLength];

	// Iterate over remaining cuts
	for reducedLength := currentLength - 1; reducedLength >= 1 ; reducedLength-- {
		maxPriceForReducedLength := RodMaximumProfitCut_Recursive(cutPrices, reducedLength);
		totalPrice := maxPriceForReducedLength + cutPrices[currentLength - reducedLength];
		currentMaxPrice = int(math.Max(float64(currentMaxPrice), float64(totalPrice)));
	}
	return currentMaxPrice;
}

// Dynamic programming - 1D
func RodMaximumProfitCut(prices []int, rodLength int) int {
	// No cut is possible if original rod length is zero
	if rodLength == 0 {
		return 0
	}

	var subProblems = make([]int, rodLength+1)
	subProblems[0] = 0 // A rod of length 0 has a price of 0

	for state := 1; state <= rodLength; state++ {
		for destination := 0; destination < state; destination++ {
			length := state - destination
			price := prices[length]

			subProblems[state] = int(math.Max(float64(subProblems[state]), float64(price+subProblems[destination])))
		}
	}

	return subProblems[rodLength]
}

// Dynamic programming - 2D
func RodMaximumProfitCutDP(prices []int, rodLength int) int {
	// No cut is possible if original rod length is zero
	if rodLength == 0 {
		return 0;
	}

	// Initialize dynamic programming table
	var grid [][]int  = make([][]int, rodLength+1);
	for i := range grid {
		grid[i] = make([]int, rodLength+1)
	}

	// Base conditions
	for i := 1; i <= rodLength; i++ {
		grid[1][i] = i * prices[1];
		grid[i][1] = prices[1];
	}

	for cut := 2; cut <= rodLength; cut++ {
		for length := 2; length <= rodLength; length++ {
			if cut > length {
				grid[cut][length] = grid[cut-1][length];
			} else {
				totalPrice := prices[cut] + grid[cut][length - cut];
				grid[cut][length] = int(math.Max(float64(totalPrice), float64(grid[cut - 1][length])));
			}
		}
	}

	return grid[rodLength][rodLength];
}





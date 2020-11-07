package dynamicProgramming

import "errors"

func subsetSum(a []int, B int) ([]int, error) {
	if a == nil || len(a) == 0 || B <=0 {
		return nil, errors.New("invalid parameters")
	}

	// Allocate DP array
	var grid [][]bool = make([][]bool, len(a) + 1)
	for i,_ := range grid {
		grid[i] = make([]bool, B+1)
	}

	// Initialize DP table - first row
	grid[0][0] = true;
	for j := 1; j <= B; j++ {
		grid[0][j] = false;
	}

	// Initialize DP table - first col
	for i := 1; i <= len(a); i++ {
		grid[i][0] = true;
	}

	// Populate grid according to S(n,B)=S(n-1,B-an) || S(n-1,B)
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= B; j++ {
			// Get current item
			 var ai int = a[i-1];    // -1 because we're accessing array 'a' and not array 'grid'

			// If j (eg sum) is less that the current item, then get the result from the cell
			// above in the previous row, else value = value excluding current number (eg, upper cell)
			// OR'd with value from previous row where column is the difference between sum and
			// value
			if j - ai < 0 {
				grid[i][j] = false || grid[i-1][j]; // grid[i-1][j - ai] = false from base case
			} else {
				grid[i][j] = grid[i-1][j-ai] || grid[i-1][j];
			}
		}
	}

	// Finding the actual values required to make the given sum
	var subset []int
	column := B;
	for i := len(a); i >= 1; i-- {
		if grid[i][column] == grid[i-1][column] {
			continue;
		}

		// Current cell different from upper cell so include number
		subset = append(subset, a[i-1])		// array is 1-based

		// column (or sum) contributed to value so subtract it and repeat
		column = column - a[i-1];            //  sum = sum - current element
	}

	// No easy way to convert List<Integer> to int[]
	return subset, nil
}

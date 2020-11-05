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


}

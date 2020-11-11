package dynamicProgramming

import "errors"

func CalculateMovement(n int) (int, error) {
	// Check invariants
	if n < 1 {
		return 0, errors.New("n must be >= 1")
	}

	// Allocate programming table
	var grid [][]int = make([][]int, n)

	// Initialize base conditions
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		grid[i][0] = 1;
		grid[0][i] = 1;
	}
}

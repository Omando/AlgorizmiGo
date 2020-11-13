package dynamicProgramming

import "errors"

func CalculateRobotMovementInGrid(n int) (int, error) {
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

	// Calculation. Fill table according to recurrence relation T[x][y] = T[x-1][y] + T[x][y-1]
	// Can either move right or down
	for x := 1; x < n; x++ {
		for y:= 1; y < n; y++ {
			grid[x][y] = grid[x-1][y] + grid[x][y-1];
		}
	}

	// Return result
	return grid[n-1][n-1], nil
}

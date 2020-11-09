package dynamicProgramming

import "errors"

func CalculateBinarySequenceSize(n int) (int, error) {
	// Check invariants
	if n < 1 {
		return 0, errors.New("n must be >= 1")
	}

	// Allocate table
	var table []int = make([]int, n+1) // cell at index 0 is unused

	// Base conditions
	table[1] = 2

	// Calculation. Fill table according to recurrence relation T[n] = 2T[n-1]
	for i := 2; i <= n; i++ {
		table[i] = 2 * table[i-1]
	}

	// Return result
	return table[n], nil
}

package recursion

import "fmt"

func Sudoku(grid [][]int) [][]int {
	// We don't want to write to the input, so deep-copy grid into result
	var result [][]int
	for i, _ := range grid {
		result[i] = make([]int, len(grid[i])) // Create a new row
		copy(result[i], grid[i])              // Initialize new row with data
	}

	// What about
	result2 := grid // value types so creates a copy
	grid[0][0] = 111
	result2[0][0] = 222
	fmt.Println(result2)

	// TODO

	return result
}

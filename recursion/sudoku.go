package recursion

func Sudoku(grid [][]int) [][]int {
	// We don't want to write to the input, so deep-copy grid into result
	// The following statement to create a copy of grid causes changes to one grid
	// to be reflected in the other
	//	var result [][]int = grid
	var result [][]int = make([][]int, len(grid))
	for i, _ := range grid {
		result[i] = make([]int, len(grid[i])) // Create a new row
		copy(result[i], grid[i])              // Initialize new row with data
	}

	// TODO

	return result
}

func canAddItem(grid [][]int, row int, column int, misisngNumber int) bool {
	// todo
	return false
}

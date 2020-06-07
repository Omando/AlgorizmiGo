package recursion

func Sudoku(grid [][]int) (bool, [][]int) {
	// Create a copy of the unsolved Sudoku puzzle. This copy will hold the actual solution
	// Note: the following statement to create a copy of grid causes changes to one grid
	// to be reflected in the other: var result [][]int = grid
	var result [][]int = make([][]int, len(grid))
	for i := range grid {
		result[i] = make([]int, len(grid[i])) // Create a new row
		copy(result[i], grid[i])              // Initialize new row with data
	}

	hasSolution := solve(grid, 0, 0, result);
	return hasSolution, result
}

func solve(grid [][]int, row int, col int, result [][]int) bool {
	// If we reached last row, move to the next column and start from the first row
	// Exit if we reached the last column
	if row == len(grid) {
		row = 0;
		col = col + 1
		if col == len(grid[0]) {
			return true
		}
	}

	// Ignore current cell if it already has a value, otherwise try to insert a legal number
	if grid[row][col] != 0 {
		solve(grid, row+1, col, result);
	} else {
		// TODO
	}

	return false;
}


func canAddItem(grid [][]int, row int, column int, number int) bool {
	// todo
	return false
}

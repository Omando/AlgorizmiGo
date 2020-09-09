package recursion

// Given a 2d grid map of 1s (land) and 0s (water), count the number of islands
// Immediate neighbours of a cell are the next adjacent horizontal, vertical,
// and diagonal cells
func IslandCount(grid [][]int) int {
	// Check if grid is empty
	if len(grid) == 0 && len(grid[0]) == 0 {
		return 0
	}

	// Iterate over the grid
	var islandCount int = 0
	for i, row := range grid {
		for j, cell := range row {
			// Ignore if cell is 'water' or is already visited (-1)
			if cell == 0 || cell == -1 {
				continue
			}

			visitCells(grid, i, j)
			islandCount++
		}
	}

	return islandCount
}

func visitCells(grid [][]int, row int, col int) {
	// Mark the current cell as visited (-1) and loop around its neighbour cells
	grid[row][col] = -1

	// Iterate over immediate neighbours (horizontal / vertical / diagonal)
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			// Check boundaries. Ignore indices outside edges
			if i == -1 || i == len(grid) || j == -1 || j == len(grid[0]) {
				continue
			}

			// Nothing to do if water or already visited
			if grid[i][j] == 0 || grid[i][j] == -1 {
				continue
			}

			// We have land. Mark it as visited and search its neighbours
			visitCells(grid, i, j)
		}
	}
}

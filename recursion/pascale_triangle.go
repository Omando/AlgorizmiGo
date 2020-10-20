package recursion

func GeneratePascalTriangle(n int) [][]int {
	var rows [][]int
	solvePascal(n, rows)
	return rows
}

func solvePascal (rowIndex int, rows [][]int) {
	// Exit condition: Add first row
	if rowIndex ==  0 {
		rows[0] = append(rows[0], 1)
		return;
	}

	// Recurse until we generate the base row at index 0
	solvePascal(rowIndex - 1, rows)

	/* Construct current row from previous row */
	// Current row length is previous row length + 1
	currentRowLength := len(rows[rowIndex-1]) + 1
	rows[rowIndex] = make([]int, currentRowLength)

	// Populate current row from previous row
	for i := 0; i < currentRowLength; i++ {
		// Left and right edges of current row are always 1
		if i == 0 || i == currentRowLength - 1 {
			rows[rowIndex][i] = 1
		} else {
			// Inner cells of current row
			rows[rowIndex][i] = rows[rowIndex-1][i-1] + rows[rowIndex-1][i]
		}
	}
}

package dynamicProgramming

/* Given a network (graph) in the form of an adjacency matrix, found the value
 of the maximum cost between two vertices, S and T. Note the following assumptions
 about the adjacency matrix:
	graph[i][j] = k: edge between vertex i and vertex j has a weight of k
	graph[i][j] = 0: there is no edge between vertex i and vertex j
*/
func CalculateMaximumCostRoute(graph [][]int) ([]int, []string) {
	var rowCount int = len(graph)
	var colCount = rowCount

	// Allocate arrays for cost and path
	var optimalCosts []int = make([]int, rowCount)
	var optimalPaths []string = make([]string, rowCount)

	// Start iterating from the last row
	for i := rowCount - 1; i >= 0; i-- {
		// For each row (vertex i), iterate over all columns (vertex j). This
		// is equivalent to iterating over all connections starting vertex i
		// Look at the table on page 29 in  DSALG_RecusionDP and not that for
		// a given row index, the column index is always >= to the row index
		// So as an optimization, we could initialize j to i
		for j := 0; j < colCount; j++ {
			// TODO
		}
	}

	return optimalCosts, optimalPaths
}

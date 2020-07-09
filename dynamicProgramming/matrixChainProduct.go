package dynamicProgramming

import "math"

// This function assumes that matrix Ai has dimensions pi-1 x pi for i = 1,2, ..., n.
// Input is an array of matrix dimensions: p0, p1, ..., pn. For example, if we have
// three matrices A1 = 3x4, A2 = 4x5 and A3 = 5x6, then the input array is {3,4,5,6}
func Calculate(dimensions []int) int {
	n := len(dimensions)

	// Auxiliary table for storing m[i,j] costs
	// m[i,j] = Minimum number of scalar multiplications needed to compute
	// the matrix A[i]A[i+1]...A[j]. For example, m[1,4] = minimum cost for
	// calculating matrix product A1A2A3A4
	// Note regarding approach to  allocating a 2D slice: One approach is to allocate
	//each slice independently; the other is to allocate a single array and point the
	//individual slices into it. If the slices might grow or shrink, they should be
	//allocated independently to avoid overwriting the next line; if not, it can be
	//more efficient to construct the object with a single allocation
	var m [][]int = make([][]int, n)      // top level slice
	var allCells []int = make([]int, n*n) // one large slice to hold all slices
	// Loop over the rows, slicing each row from the front of the remaining allCells slice
	for i := range m {
		m[i] = allCells[:n]
		allCells = allCells[n:]
	}

	var j, min, q int
	// Refer to RecursionDP document and note the colored squares to
	// explain the use of diff
	for diff := 1; diff < n-1; diff++ {
		for i := 1; i < n-diff; i++ {
			j = diff + i // see RecursionDP document
			min = math.MaxInt32

			// Actual computation of the minimum cost
			// Note that k start from i: m[1][1] + m[2][3] + ...
			//                           m[2][2] + m[3][4] + ....
			for k := i; k <= j-1; k++ {
				q = m[i][k] + m[k+1][j] + (dimensions[i-1] * dimensions[k] * dimensions[j])
				if q < min {
					min = q
				}
			}
			m[i][j] = min
		}
	}
	return m[1][n-1] // Top right corner (excluding row zero)
}

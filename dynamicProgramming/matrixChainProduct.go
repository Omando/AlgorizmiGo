package dynamicProgramming

// This function assumes that matrix Ai has dimensions pi-1 x pi for i = 1,2, ..., n.
// Input is an array of matrix dimensions: p0, p1, ..., pn. For example, an input array
// of size 4 such as {3,4,5,6} mean we have 3 matrices as follows:
//  A1 = 3x4
//  A2 = 4x5
//  A3 = 5x6
func Calculate(dims []int) int {
	n := len(dims)

	// Auxiliary table for storing c[i,k] costs
	// c[i,k] = Minimum number of scalar multiplications needed to compute the matrix
	// A[i]A[i+1]...A[k]. For example, m[1,4] is minimum cost for calculating matrix
	// product A1A2A3A4. Because the table is initialized to zeros, base conditions
	// are also initialized (grid[i][i] = 0)
	// Recall that if dim.length is n then there are n-1 matrices
	var C [][]int = make([][]int, n)
	var allCells []int = make([]int, n*n)

	// Allocating 2D array: loop over the rows, slicing each row from the front of the
	// remaining allCells slice
	for i,_ := range C {
		C[i] = allCells[:n]
		allCells = allCells[n:]
	}


	// C(i,k)= C(i,j) + C(j+1,k) + RiCjCk
	for i := 1; i < n; i++ {       // We ignore row 0 and column 0
		//C(1,2) = C(1,1) + C(2,2) + r1c1c2
		//C(2,3) = C(2,2) + C(3,3) + r2c2c3
		//C(3,4) = C(3,3) + C(4,4) + r3c3c4
		//C(1,3) = min(C(1,1) + C(2,3) + r1c1c3, C(1,2) + C(3,3) + r1c2c3)

	}
	return C[1][n-1];       // we ignore row 0 and col 0
}


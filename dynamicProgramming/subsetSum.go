package dynamicProgramming

import "errors"

func subsetSum(a []int, B int) ([]int, error) {
	if a == nil || len(a) == 0 || B <=0 {
		return nil, errors.New("invalid parameters")
	}

	// Allocate DP array
	var grid [][]bool = make([][]bool, len(a) + 1)
	for i,_ := range grid {
		grid[i] = make([]bool, B+1)
	}




}

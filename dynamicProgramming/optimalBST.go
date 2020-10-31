package dynamicProgramming

import "errors"

func FindOptimalBST(keys []byte, freq []int) (costs [][]int, roots [][]int, err error) {
	// keys and freq table should bath have the same length
	if len(keys) != len(freq) {
		return nil, nil, errors.New("keys and freq arrays should have equal lengths");
	}

	// Allocate one programming table for optimal costs, and another for optimal root nodes
	length :=  len(keys)
	var costs [][]int = make([][]int, length)
	var roots [][]int = make([][]int, length)
	for i := range keys {
		costs[i] = make([]int, length)
		roots[i] = make([]int, length)
	}

	// Pre populate costs and roots tables
	for i := 0; i < length; i++ {
		costs[i][i] = freq[i]	// populate diagonal data
		roots[i][i] = i			// populate diagonal data
 	}

	// Calculate optimal costs (diagonally, see background colors in document)
	offset := 1
	for offset < length {
		i := 0;

		//  i <= r <= j
		for j := offset; j < length; j++ {
			// todo
		}

		offset++;   // Increase column offset for next diagonal
	}
	
	// todo
	return costs, roots, nil
}

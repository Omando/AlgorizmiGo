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
		costs[i][i] = freq[i]
		roots[i][i] = i
 	}
 	





	// todo
	return costs, roots, nil
}

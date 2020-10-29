package dynamicProgramming

import "errors"

func FindOptimalBST(keys []byte, freq []int) (costs [][]int, roots [][]int, err error) {
	// keys and freq table should bath have the same length
	if len(keys) != len(freq) {
		return nil, nil, errors.New("keys and freq arrays should have equal lengths");
	}
	
	// todo
	return costs, roots, nil
}

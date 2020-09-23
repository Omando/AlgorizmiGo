package recursion

import (
	"AlgorizmiGo/recursion"
	"testing"
)

func TestSudoku(t *testing.T) {
	tests := []struct {
		input            [][]int
		expectedSolution [][]int
	}{
		{input: getInput1(), expectedSolution: getExpectedSolution1()},
	}

	for _, test := range tests {
		recursion.Sudoku(test.input)
	}
}

func getInput1() [][]int {
	return [][]int{
		{5, 3, 1, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
}

func getExpectedSolution1() [][]int {
	return [][]int{}
}

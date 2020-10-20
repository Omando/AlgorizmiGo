package recursion

import (
	"AlgorizmiGo/recursion"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	tests := [] struct {
		rowsIndex    int
		expectedRows [][]int
	} {
		{rowsIndex: 0, expectedRows: getExpectedRows(0)},
		{rowsIndex: 1, expectedRows: getExpectedRows(1)},
		{rowsIndex: 2, expectedRows: getExpectedRows(2)},
		{rowsIndex: 5, expectedRows: getExpectedRows(5)},
		{rowsIndex: 6, expectedRows: getExpectedRows(6)},
	}

	for _, test := range tests {
		actualRows := recursion.GeneratePascalTriangle(test.rowsIndex)

		for i, row := range test.expectedRows {
			assert.EqualValues(t, row, actualRows[i])
		}
	}
}

func getExpectedRows(rowsIndex int) [][]int {
	if rowsIndex > 6 {
		errors.New("A maximum of 6 test rows is supported")
	}

	rows := [][]int {
		[]int{1},
		[]int {1,1},
		[]int {1,2,1},
		[]int {1,3,3,1},
		[]int {1,4,6,4,1},
		[]int {1,5,10,10,5,1},
		[]int {1,6,15,20,15,6,1},
	}

	// Return only the requested number of test rows
	return rows[:rowsIndex+1]
}


package recursion

import (
	"AlgorizmiGo/recursion"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	tests := []struct {
		numberOfRows int
		expectedRows [][]int
	}{
		{numberOfRows: 1, expectedRows: getExpectedRows(1)},
	}

	for _, test := range tests {
		actualRows := recursion.GeneratePascalTriangle(test.numberOfRows)

		for i, row := range test.expectedRows {
			assert.EqualValues(t, row, actualRows[i])
		}
	}
}

func getExpectedRows(numberOfRows int) [][]int {
	if numberOfRows > 6 {
		errors.New("A maximum of 6 test rows is supported")
	}

	rows := [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
		[]int{1, 5, 10, 10, 5, 1},
		[]int{1, 6, 15, 20, 16, 6, 1},
	}

	// Return only the requested number of test rows
	return rows[:numberOfRows]
}

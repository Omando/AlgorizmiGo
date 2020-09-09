package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIslandsCount(t *testing.T) {
	tests := []struct {
		grid          [][]int
		expectedCount int
	}{
		{grid: getGridWith4Islands(), expectedCount: 4},
		{grid: getGridWith3Islands(), expectedCount: 3},
	}

	for _, test := range tests {
		actualCount := recursion.IslandCount(test.grid)
		assert.Equal(t, test.expectedCount, actualCount)
	}
}

func getGridWith4Islands() [][]int {
	return [][]int{
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1}}
}

func getGridWith3Islands() [][]int {
	return [][]int{
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 1}}
}

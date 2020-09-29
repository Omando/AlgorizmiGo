package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackUnwind(t *testing.T) {
	recursion.StackUnwind(5)
}
func TestRecursionWithDataCaches(t *testing.T) {
	recursion.RecursionWithDataCaches(5)
}

func TestForwardIterationViaRecursion(t *testing.T) {
	inputs := []struct {
		start, end, increment int
	}{
		{0, 10, 1},
		{0, 10, 2},
		{0, 10, 5},
		{0, 10, 10},
	}

	for _, input := range inputs {
		recursion.ForwardIterationViaRecursion(input.start, input.end, input.increment)
	}
}

func TestReverseIterationViaRecursion(t *testing.T) {
	inputs := []struct {
		start, end, decrement int
	}{
		{10, 0, 1},
		{10, 0, 2},
		{10, 0, 5},
		{10, 0, 10},
	}

	for _, input := range inputs {
		recursion.ReverseIterationViaRecursion(input.start, input.end, input.decrement)
	}
}

func TestGridIterationViaRecursion(t *testing.T) {
	inputs := []struct {
		grid [][]int
	}{
		{grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
	}

	for _, input := range inputs {
		actualOutput := recursion.GridIterationViaRecursion(input.grid)

		for i := 0; i < len(input.grid); i++ {
			assert.ElementsMatch(t, input.grid[i], actualOutput[i])
		}
	}
}

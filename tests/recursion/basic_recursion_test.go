package recursion

import (
	"AlgorizmiGo/recursion"
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

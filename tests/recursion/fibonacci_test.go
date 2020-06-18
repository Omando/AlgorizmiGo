package recursion

import (
	"AlgorizmiGo/recursion"
	"fmt"
	"testing"
)

// Test data
var testData = []struct {
	n        int
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestShouldCalculateFibonacci_UsingIteration(t *testing.T) {
	for _, input := range testData {
		fmt.Printf("Calculating fibonacii for %v\n", input.n)
		var actualResult = recursion.FibonacciIterative(input.n)

		if actualResult != input.expected {
			t.Errorf("Expected %v but got %v", input.expected, actualResult)
		}
	}
}

func TestShouldCalculateFibonacci_UsingRecursion(t *testing.T) {
	for _, input := range testData {
		fmt.Printf("Calculating fibonacii for %v\n", input.n)
		var actualResult = recursion.FibonacciRecursive(input.n)

		if actualResult != input.expected {
			t.Errorf("Expected %v but got %v", input.expected, actualResult)
		}
	}
}

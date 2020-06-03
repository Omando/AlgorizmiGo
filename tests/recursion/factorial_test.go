package recursion

import (
	"AlgorizmiGo/recursion"
	"fmt"
	"testing"
)

func TestShouldCalculateFactorial(t *testing.T) {
	fmt.Println("Testing individual input")
	var result, err = recursion.Factorial(5)

	if err != nil {
		t.Errorf("No error should have been thrown ")
	}

	if result != 120 {
		t.Errorf("Expected 120 but got %v", result)
	}
}

func TestCalculateMultipleFactorials(t *testing.T) {
	fmt.Println("Testing multiple inputs")
	inputs := []struct {
		n              int
		expectedResult int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, input := range inputs {
		fmt.Printf("Calculating factorial for %v\n", input.n)
		var actualResult, err = recursion.Factorial(input.n)
		if err != nil {
			t.Errorf("No error should have been thrown ")
		}

		if actualResult != input.expectedResult {
			t.Errorf("Expected %v but got %v", input.expectedResult, actualResult)
		}
	}
}

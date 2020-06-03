package recursion

import "errors"

func Factorial(n int) (int, error) {
	// Exit conditions
	if n < 0 {
		return n, errors.New("input must be greater or equal to zero")
	}

	if n == 0 || n == 1 {
		return 1, nil
	}

	// Recursion
	nMinus1Factorial, err := Factorial(n - 1)
	return n * nMinus1Factorial, err
}

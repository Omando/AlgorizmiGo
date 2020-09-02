package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreatestCommonDivisor(t *testing.T) {
	tests := []struct {
		n           int
		m           int
		expectedGCD int
	}{
		{84, 56, 28},
		{56, 84, 28},
		{86, 56, 2},
		{56, 86, 2},
		{32, 21, 1},
		{21, 32, 1},
		{1008, 21, 21},
		{21, 1008, 21},
		{3, 6, 3},
		{6, 3, 3},
		{12345, 765, 15},
		{765, 12345, 15},
		{765, 765, 765},
	}

	for _, test := range tests {
		actualGCD := recursion.GCD(test.n, test.m)
		assert.Equal(t, test.expectedGCD, actualGCD)
	}
}

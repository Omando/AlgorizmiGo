package recursion

import (
	"AlgorizmiGo/recursion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		stepsCount int
		expectedNumberOfWays int
	} {
		{stepsCount: 1, expectedNumberOfWays: 1},
		{2, 2, },
		{3, 3, },
		{4, 5, },
		{5, 8, },
		{6, 13, },
		{7, 21, },
	}
	for _, test := range tests {
		actualCountOfWays := recursion.ClimbStairs(test.stepsCount)
		assert.Equal(t, test.expectedNumberOfWays, actualCountOfWays)
	}
}

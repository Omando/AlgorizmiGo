package dynamicProgamming

import (
	"AlgorizmiGo/dynamicProgramming"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_calculate_grid_robot_movement(t *testing.T) {
	tests := []struct {
		n            int
		expectedCount int
	}{
		{n: 1, expectedCount: 1},
		{n: 2, expectedCount: 2},
		{n: 3, expectedCount: 6},
		{n: 4, expectedCount: 20},
		{n: 5, expectedCount: 70},
	}

	for _, test := range tests {
		actualCount, err := dynamicProgramming.CalculateRobotMovementInGrid(test.n)
		assert.Nil(t, err)
		assert.Equal(t, test.expectedCount, actualCount)
	}
}

package dynamicProgamming

import (
	"AlgorizmiGo/dynamicProgramming"
	"testing"
)

// Generated using CTRL+SHIFT+T
func TestCalculateMaximumCostRoute(t *testing.T) {

	basicGraph := [][]int{
		{0, 10, 11, 12, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 11, 10, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 9, 13, 12, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 12, 13, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 13, 12, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 14, 10, 11, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 14, 12, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	tests := []struct {
		name            string
		graph           [][]int
		expectedMaxCost int
		expectedMaxPath string
	}{
		{name: "Basic", graph: basicGraph, expectedMaxCost: 50, expectedMaxPath: "0 --> 2 --> 5 --> 7 --> 10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualMaxCost, actualMaxPath := dynamicProgramming.CalculateMaximumCostRoute(tt.graph)
			if actualMaxCost != tt.expectedMaxCost {
				t.Errorf("CalculateMaximumCostRoute() actual max cost:%v, expected:%v", actualMaxCost, tt.expectedMaxCost)
			}

			if actualMaxPath != tt.expectedMaxPath {
				t.Errorf("CalculateMaximumCostRoute() actual max path:%v, expected:%v", actualMaxPath, tt.expectedMaxPath)
			}
		})
	}
}

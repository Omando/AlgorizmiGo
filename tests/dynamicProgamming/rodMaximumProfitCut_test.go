package dynamicProgamming

import (
	"AlgorizmiGo/dynamicProgramming"
	"testing"
)

func TestRodMaximumProfitCut(t *testing.T) {
	tests := []struct {
		name      string
		prices    []int
		rodLength int
		want      int
	}{
		{"rod length 4", []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}, 4, 10},
		{"rod length 5", []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}, 5, 13},
		{"rod length 6", []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}, 6, 17},
		{"rod length 7", []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}, 7, 18},
		{"rod length 8", []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}, 8, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test recursive
			if got := dynamicProgramming.RodMaximumProfitCut_Recursive(tt.prices, tt.rodLength); got != tt.want {
				t.Errorf("RodMaximumProfitCut_Recursive() = %v, want %v", got, tt.want)
			}

			// Test DP
			if got := dynamicProgramming.RodMaximumProfitCutDP(tt.prices, tt.rodLength); got != tt.want {
				t.Errorf("RodMaximumProfitCutDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

package dynamicProgamming

import (
	"AlgorizmiGo/dynamicProgramming"
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		dimensions []int
	}

	tests := []struct {
		name  string
		input args
		want  int
	}{
		{name: "4 matrices", input: args{[]int{5, 4, 6, 2, 7}}, want: 158},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dynamicProgramming.Calculate(tt.input.dimensions); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

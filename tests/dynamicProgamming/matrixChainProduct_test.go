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
		{name: "2 matrices", input: args{[]int{5, 4, 6}}, want: 120},
		{name: "3 matrices", input: args{[]int{5, 4, 6, 2}}, want: 88},
		{name: "4 matrices", input: args{[]int{5, 4, 6, 2, 7}}, want: 158},
		{name: "5 matrices", input: args{[]int{10, 20, 30, 40, 30}}, want: 30000},
		{name: "5 matrices", input: args{[]int {4,10,3,12,20,7,6,20,13}}, want: 2682},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dynamicProgramming.Calculate(tt.input.dimensions); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

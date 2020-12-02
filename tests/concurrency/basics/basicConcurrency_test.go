package concurrency_basics

import (
	"AlgorizmiGo/concurrency/basics"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_calculate_sum_concurrently(t *testing.T) {
	tests := []struct {
		input 		[]int
		expectedSum	int
	} {
		{[]int{1,2,3,4}, 10},
		{[]int{-1,1,-10,10}, 0},
		{[]int{2,4,6,8,10}, 30},
	}

	for _, test := range tests {
		actualSum := concurrency_basics.ConcurrentSum(test.input)
		assert.EqualValues(t, test.expectedSum, actualSum)
	}
}

func Test_should_receive_until_closed(t *testing.T) {
	tests := []struct {
		input 		int
		expectedSum	int
	} {
		{5, 15},
		{6, 21},
	}

	for _, test := range tests {
		actualSum := concurrency_basics.ReceiveUntilClose(test.input)
		assert.EqualValues(t, test.expectedSum, actualSum)
	}
}

func Test_produce_and_consumer_values(t *testing.T) {
	tests := []struct {
		valueCount 		int
		generator		concurrency_basics.ValueGenerator
		expectedOutput	[]int
	} {
		{5,  concurrency_basics.GenerateIncrementalValues,[]int{1,2,3,4,5}},
		{7,  concurrency_basics.GenerateIncrementalValues,[]int{1,2,3,4,5,6,7}},
		{1,  concurrency_basics.GenerateIncrementalValues,[]int{1
		}},
	}

	for _, test := range tests {
		actualOutput := concurrency_basics.BasicProducerConsumer(test.generator, test.valueCount)
		assert.EqualValues(t, test.expectedOutput, actualOutput)
	}
}
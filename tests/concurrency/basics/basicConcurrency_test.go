package concurrency_basics

import (
	. "AlgorizmiGo/concurrency/basics"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
	"time"
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
		actualSum := ConcurrentSum(test.input)
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
		actualSum := ReceiveUntilClose(test.input)
		assert.EqualValues(t, test.expectedSum, actualSum)
	}
}

func Test_produce_and_consumer_values(t *testing.T) {
	tests := []struct {
		valueCount     int
		generator      ValueGenerator
		expectedOutput []int
	} {
		{5, GenerateIncrementalValues,[]int{1,2,3,4,5}},
		{7, GenerateIncrementalValues,[]int{1,2,3,4,5,6,7}},
		{1, GenerateIncrementalValues,[]int{1}},
	}

	for _, test := range tests {
		actualOutput := BasicProducerConsumer(test.generator, test.valueCount)
		assert.EqualValues(t, test.expectedOutput, actualOutput)
	}
}

func Test_atomic_counter_should_count_when_used_in_multiple_threads(t *testing.T) {
	tests := []struct {
		threadCount        int
		countPerThread     int
		expectedFinalCount int
	} {
		{1, 100, 100},
		{5, 100, 500},
		{10, 100, 1000},
		{20, 1000, 20000},
	}

	for _, test := range tests {
	
	}
}
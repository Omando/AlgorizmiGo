package recursion

import (
	"AlgorizmiGo/collections/sequential"
	"fmt"
)

func StackUnwind(n int) {
	// Exit condition
	if n < 0 {
		return
	}

	fmt.Printf("Before recursion - %v\n", n)
	StackUnwind(n - 1)
	fmt.Printf("After recursion - %v\n", n)
}

var preQueue sequential.Queue = sequential.CreateQueue()
var postQueue sequential.Queue = sequential.CreateQueue()
var preStack sequential.Stack = sequential.CreateStack(5)
var postStack sequential.Stack = sequential.CreateStack(5)

func RecursionWithDataCaches(n int) {
	doRecursionWithDataCaches(n)
	dumpRecursionCaches()
}

func doRecursionWithDataCaches(n int) {
	// Exit condition
	if n == 0 {
		return
	}

	// Cache data before recursion
	preQueue.Enqueue(n)
	preStack.Push(sequential.Item{n})

	doRecursionWithDataCaches(n - 1)

	// Cache data after recursion
	postQueue.Enqueue(n)
	postStack.Push(sequential.Item{n})
}

func dumpRecursionCaches() {
	fmt.Println("PreStack:")
	for !preStack.IsEmpty() {
		value, err := preStack.Pop()
		if err == nil {
			fmt.Printf("%v ", value)
		}
	}

	fmt.Println("\nPreQueue:")
	for !preQueue.IsEmpty() {
		value, err := preQueue.Dequeue()
		if err == nil {
			fmt.Printf("%v ", value)
		}
	}

	fmt.Println("\nPostStack:")
	for !postStack.IsEmpty() {
		value, err := postStack.Pop()
		if err == nil {
			fmt.Printf("%v ", value)
		}
	}

	fmt.Println("\nPostQueue:")
	for !postQueue.IsEmpty() {
		value, err := postQueue.Dequeue()
		if err == nil {
			fmt.Printf("%v ", value)
		}
	}
}

type incrementer func(index int) int

func ForwardIterationViaRecursion(start int, end int, increment int) {
	if start > end {
		return
	}

	var funcIncrementer incrementer = func(index int) int { return index + increment }
	doForwardIteration(start, end, funcIncrementer)
}

func doForwardIteration(index int, end int, funcIncrementer incrementer) {
	// Exit condition
	if index >= end {
		return
	}

	fmt.Println("Processing index ", index)
	doForwardIteration(funcIncrementer(index), end, funcIncrementer)
}

func ReverseIterationViaRecursion(start int, end int, decrement int) {
	if start < end {
		return
	}

	var funcDecrementer incrementer = func(index int) int { return index - decrement }
	doReverseIteration(start, end, funcDecrementer)
}

func doReverseIteration(index int, end int, funcDecrementer incrementer) {
	// Exit condition
	if index <= end {
		return
	}

	fmt.Println("Processing index ", index)
	doForwardIteration(funcDecrementer(index), end, funcDecrementer)
}

func GridIterationViaRecursion(grid [][]int) [][]int {
	var result [][]int = make([][]int, len(grid))
	for i := range grid {
		result[i] = make([]int, len(grid[0]))
	}
	return doGridIterationViaRecusion(grid, 0, 0, result)
}

func doGridIterationViaRecusion(grid [][]int, row int, col int, result [][]int) [][]int {
	// If we reached last row, move to the next column and start from the first row
	// Exit if we reached the last column
	if row == len(grid) {
		row = 0;
		col = col + 1
		if col == len(grid[0]) {
			return result
		}
	}

	result[row][col] = grid[row][col]
	return doGridIterationViaRecusion(grid, row + 1, col, result)
}

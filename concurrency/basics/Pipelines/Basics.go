package Pipelines

import "fmt"

 /* Pipelines 101
 This code is VERY basic and minimal. It only serves to illustrate the overall concept of
 pipelines using channels
 Output:
	Receive value: 0
	Receive value: 1
	Receive value: 4
	Receive value: 9
	Receive value: 16
	Receive value: 25
	Receive value: 36
	Receive value: 49
	Receive value: 64
	Receive value: 81
 */
func Run() {
	// Create channel used to transfer information through the pipeline
	numbers := make(chan int)
	squares := make(chan int)

	// The pipeline is composed of 3 subroutines with each piece of information
	// (a number in this case) flowing from counter to squarer to printer.
	go counter(numbers)					// publish on numbers
	go squarer(numbers, squares)		// receive on numbers and publish on squares
	go printer(squares)					// receive on squares
}

// Pipeline stage 1: generate number 0from 0 to 9 on channel ch
func counter(ch chan int)  {
	for x:=0; x < 10; x++ {
		ch <- x
	}
}

// Pipeline stage 2: Receive numbers from channel, square them, then send results
// on a different channel
func squarer(receive chan int, send chan int) int {
	for {
		value := <- receive
		send <- value * value
	}
}

// Pipeline stage 2: Receive results from channel and display
func printer(receive chan int) {
	for {
		fmt.Printf("Receive value: %d\n", <- receive)
	}
}

package Pipelines

import (
	"fmt"
)

/* Pipelines 102
range is used to receive all the values sent on a channel and then terminate
the loop after the last one. Also unidirectional channels are used to restrict
a channel to send-only or receive-only
*/
func RunWithRange() {
	// Create channel used to transfer information through the pipeline
	numbers := make(chan int)
	squares := make(chan int)

	// The pipeline is composed of 3 subroutines with each piece of information
	// (a number in this case) flowing from counter to squarer to printer.
	go counterv2(numbers)          // publish on numbers
	go squarerv2(numbers, squares) // receive on numbers and publish on squares
	go printerv2(squares)          // receive on squares
}

// Pipeline stage 1: generate number 0from 0 to 9 on channel ch
func counterv2(outgoing chan<- int) {
	defer close(outgoing)
	for x := 0; x < 10; x++ {
		outgoing <- x
	}
}

// Pipeline stage 2: Receive numbers from channel, square them, then send results
// on a different channel
func squarerv2(incoming <-chan int, outgoing chan<- int) {
	defer close(outgoing)

	// Note use of range to receive values and detect when a channel is closed
	for value := range incoming {
		outgoing <- value * value
	}
}

// Pipeline stage 2: Receive results from channel and display
func printerv2(incoming <-chan int) {

	// Note use of range to receive values and detect when a channel is closed
	for value := range incoming {
		fmt.Printf("Receive value: %d\n", value)
	}
}

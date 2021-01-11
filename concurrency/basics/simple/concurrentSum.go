package simple

//  Sum the contents of a slice and send results through channel
func sum(input []int, ch chan int) {
	// Calculate sum
	sum := 0
	for _, value := range input {
		sum += value
	}

	// Transmit value through channel
	ch <- sum
}

func ConcurrentSum(input []int) int {
	// Create a communication channel for sending and receiving sums
	ch := make(chan int)

	// Invoke two goroutines, first one operates on the first half of the slice, and the
	// second one operates on the other half
	halfLength := len(input)/2
	go sum(input[:halfLength], ch)
	go sum(input[halfLength:], ch)

	// Wait for results
	leftSum := <- ch
	rightSum := <- ch

	return leftSum + rightSum
}

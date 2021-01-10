package simple

func ReceiveUntilClose(upperLimit int) int {
	// Create an unbuffered channel
	ch := make(chan int)
	go startSender(ch, upperLimit)

	// Receive values until the channel is closed
	var sum int = 0
	for txValue := range ch {
		sum += txValue
	}

	return sum
}

func startSender(ch chan int, upperLimit int) {
	// Send itemCount values
	for i := 0; i <= upperLimit; i++ {
		ch <- i
	}

	close(ch)
}

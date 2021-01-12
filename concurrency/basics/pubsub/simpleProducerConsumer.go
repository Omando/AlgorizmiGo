package pubsub

import (
	"sync"
)

var wg sync.WaitGroup
type ValueGenerator func(previousValue int) int

func BasicProducerConsumer(valueGenerator ValueGenerator, valueCount int) (results []int) {
	wg.Add(2); // Waiting on 2 go routines

	// Create channels to communicate results and feedback
	resultsChannel := make(chan int)
	feedbackChannel := make(chan string)

	// Run producer and consumer
	go runProducer(resultsChannel, feedbackChannel, valueGenerator)
	go runConsumer(resultsChannel, feedbackChannel, valueCount, &results)

	// Wait for go routines to complete
	wg.Wait()
	return results
}

// Consume some results then tell producer to stopC
func runConsumer(resultsChannel chan int, feedbackChannel chan string, valueCount int, result *[]int) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	// Consumer 5 values then tell producer to stop
	for i := 0; i < valueCount; i++ {
		value := <- resultsChannel
		*result = append(*result, value)
	}

	// Tell consumer to stop
	feedbackChannel <- "quit"
}

// Produce values until told to stop
func runProducer(resultsChannel chan int, feedbackChannel chan string, valueGenerator ValueGenerator) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	value := valueGenerator(0)
	for {			// while loop
		select {
			case resultsChannel <- value:
				value = valueGenerator(value)			// Generate next value
			case  feedback := <- feedbackChannel:
				if feedback == "quit" {
					close(feedbackChannel)
					close(resultsChannel)
					return
				}
		}
	}
}

func GenerateIncrementalValues(previousValue int) int {
	return previousValue + 1
}

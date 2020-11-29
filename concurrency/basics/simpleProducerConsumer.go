package concurrency_basics

import (
	"sync"
)

var wg sync.WaitGroup
type ValueGenerator func(previousValue int) int

func BasicProducerConsumer(valueGenerator ValueGenerator, valueCount int) (results []int) {
	wg.Add(2);

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

// Consume some results then tell producer to stop
func runConsumer(resultsChannel chan int, feedbackChannel chan string, valueCount int, result *[]int) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()


}

// Produce values until told to stop
func runProducer(resultsChannel chan int, feedbackChannel chan string, valueGenerator ValueGenerator) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()


}

func GenerateIncrementalValues(previousValue int) int {
	return previousValue + 1
}

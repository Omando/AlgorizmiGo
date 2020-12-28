package Pipelines

import (
	"fmt"
	"sync"
)

func RunConcurrentStages() {
	// Create channels used to transfer information through the pipeline
	numbersCh := make(chan int)

	/* The pipeline is composed of 3 stages */
	// Stage1: publish on numbersCh channel
	go counterV3(numbersCh)

	// Stage2: Receive on multiple numbersCh channels, then merge results and publish
	// merge results on mergedSquaresCh channel
	c1 := squarerV3(numbersCh)
	c2 := squarerV3(numbersCh)
	mergedSquaresCh := squarerMerger(c1, c2)

	// Stage3: receive on mergedSquaresCh channel
	go printerV3(mergedSquaresCh)

	var dummy string
	fmt.Scanln(&dummy)
}

// Pipeline stage 1: generate numbers from 0 to 9 on outgoing channel
func counterV3(outgoing chan <- int) {
	defer close(outgoing)
	for x := 0; x < 10; x++ {
		fmt.Printf("Generating item %d\n", x)
		outgoing <- x
	}
}

// Pipeline stage 2: Receive numbers from inboundNumbers channel, square them, then send
// results on an outgoing channel
func squarerV3(inboundNumbers <- chan int) <- chan int {
	outBound := make(chan int)
	go func() {
		for data := range inboundNumbers {
			fmt.Printf("Calculating square for item  %d\n", data)
			outBound <- data * data
		}
	}()
	return outBound
}

// The squarerMerger function converts a list of channels to a single channel by starting
//a goroutine for each inbound channel that copies the values to the sole outbound channel.
//Once all the output goroutines have been started, merge starts one more goroutine to close
//the outbound channel after all sends on that channel are done.
func squarerMerger(inboundChannels... <- chan int) <- chan int {
	// Setup wait group to wait for squarer goroutines
	var wg sync.WaitGroup;
	wg.Add(2)

	// Create a channel to publish merge results
	out := make(chan int)

	// squarerV3 publishes on inboundChannels, and squarerMerger receives results
	// by listening on inboundChannels. All results received from inboundChannels
	// are merged by publishing them on a dedicated output channel
	for _, inboundChannel := range inboundChannels {
		go func(c <- chan int) {
			for n := range c {
				fmt.Printf("Merging item %d\n", n)
				out <- n
			}
			wg.Done()
		}(inboundChannel)
	}

	// Start a goroutine to close out once all output goroutines are done
	go func() {
		wg.Wait()
		fmt.Printf("Merging completed\n")
		close(out)
	}()

	return out
}

// Pipeline stage 3: Receive results from incoming channel and display
func printerV3(incoming <-chan int) {

	// Note use of range to receive values and detect when a channel is closed
	for value := range incoming {
		fmt.Printf("Receive value: %d\n", value)
	}
}


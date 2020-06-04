package collections

import (
	"AlgorizmiGo/collections/concurrent" // note use of fully qualified path
	"container/list"
	"fmt"
	"github.com/cucumber/godog"
	"os"
	"sync"
	"testing"
)

type BoundedQueueTest struct {
	consumerCount int
	producerCount int
	capacity      int
	buffer        *concurrent.BoundedBuffer
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
		//Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	}

	// godog v0.9.0 (latest) and earlier
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opts)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

var boundedQueueTest BoundedQueueTest

func thereAreProducers(producerCount int) error {
	boundedQueueTest.producerCount = producerCount
	return nil
}

func thereAreConsumers(consumerCount int) error {
	boundedQueueTest.consumerCount = consumerCount
	return nil
}

func boundedQueueCapacity(capacity int) error {
	boundedQueueTest.capacity = capacity
	return nil
}

func whenProducersAndConsumersRun() error {
	boundedQueueTest.buffer = &concurrent.BoundedBuffer{
		Data:      list.New(),
		Capacity:  boundedQueueTest.capacity,
		Condition: sync.NewCond(&sync.Mutex{}),
	}

	// Two waitgroups are used to synchronize the start and end of goroutines.
	// All goroutines wait on startGate to be opened by this function
	var startGate sync.WaitGroup
	startGate.Add(1)

	// This function waits on this endGate to be opened once the last goroutine completes
	// producers
	var endGate sync.WaitGroup
	endGate.Add(boundedQueueTest.consumerCount + boundedQueueTest.producerCount)

	for i := 0; i < boundedQueueTest.producerCount; i++ {
		go func(itemToAdd int) {
			startGate.Wait()
			boundedQueueTest.buffer.Enqueue(itemToAdd)
			endGate.Done()
			fmt.Println("Adding completed")
		}(i)
	}

	// consumers
	for i := 0; i < boundedQueueTest.consumerCount; i++ {
		go func() {
			startGate.Wait()
			item := boundedQueueTest.buffer.Dequeue()
			endGate.Done()
			fmt.Println("Removed ", item)
		}()
	}

	// Start all goroutines
	fmt.Println("Starting all goroutines")
	startGate.Done()

	// Wait until all go routines are done
	endGate.Wait()
	fmt.Println("All goroutines completed")
	return nil
}

func thereShouldBeItemsRemaining(remainingItemCount int) error {
	actualItemCount := boundedQueueTest.buffer.Data.Len()
	if remainingItemCount != actualItemCount {
		return fmt.Errorf("Expected %d items remaining, but found %d",
			remainingItemCount, actualItemCount)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there are (\d+) producers$`, thereAreProducers)
	s.Step(`^there are (\d+) consumers$`, thereAreConsumers)
	s.Step(`^bounded queue capacity is (\d+)$`, boundedQueueCapacity)
	s.Step(`^when producers and consumers run$`, whenProducersAndConsumersRun)
	s.Step(`^there should be (\d+) items remaining$`, thereShouldBeItemsRemaining)
}

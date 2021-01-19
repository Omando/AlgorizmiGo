package diretoryTraversal

import (
	"sync"
	"time"
)

// TraversRoots walks all roots and reports result when all roots have been processed
func TraverseRoots(roots [] string) {
	fileSizes := make(chan int64)

	// Traverse the file tree
	go func() {
		for _, dir := range roots {
			walkDir(dir, fileSizes)
		}

		// Completed processing all roots, so close channel
		close(fileSizes)
	}()

	// Execution arrives here immediately while the previous goroutine gets scheduled
	// However, read from fileSizes blocks until the first result is available
	var numberOfFiles int32
	var totalFileSizes int64
	for fileSize := range  fileSizes {
		numberOfFiles++
		totalFileSizes += fileSize
	}

	printDiskUsage(numberOfFiles, totalFileSizes)
}

// TraverseRootsAndMultiplexResults walks all roots and reports results every n milliseconds
func TraverseRootsAndMultiplexResults(roots [] string) {

	// Channel to collect results from walking a directory
	fileSizes := make(chan int64)

	// Channel to periodically report results
	var timer <-chan time.Time = time.Tick(500 * time.Microsecond)

	// Traverse the file tree
	go func() {
		for _, dir := range roots {
			walkDir(dir, fileSizes)
		}

		// Completed processing all roots, so close channel
		close(fileSizes)
	}()

	var numberOfFiles int32
	var totalFileSizes int64

loop:
	for {
		select {
		// test whether the fileSizes channel has been closed, using the two-result
		// form of receive operation.
		case fileSize, ok := <- fileSizes:
			if !ok {
				// Code arrive here when callingL close(fileSizes)
				// If the channel has been closed, the program breaks out of the loop. The labeled
				// break statement breaks out of both the select and the for loop; an unlabeled break
				// would break out of only the select, causing the loop to begin  the next iteration.

				// Print file totals. This is required because processing may complete and fileSizes
				// channel closes, but timer has yet some time to elapse
				printDiskUsage(numberOfFiles, totalFileSizes)
				break loop
			}

			// channel not closed yet, so update state
			numberOfFiles++
			totalFileSizes += fileSize
		case <- timer:
			printDiskUsage(numberOfFiles, totalFileSizes)
		}
	}
}

// TraverseRootsConcurrentlyAndMultiplexResults walks all roots concurrently and reports results
//every n milliseconds
func TraverseRootsConcurrentlyAndMultiplexResults(roots [] string) {

	// Channel to collect results from walking a directory
	fileSizes := make(chan int64)

	// Channel to periodically report results
	var timer <-chan time.Time = time.Tick(500 * time.Microsecond)

	// Wait group to synchronize dirWalk
	var wg sync.WaitGroup

	// Traverse the file tree
	go func() {
		for _, dir := range roots {
			wg.Add(1)
			walkDir(dir, fileSizes)
		}
	}()

	go func() {
		// Wait until all dirWalks have completed, then close channel
		wg.Wait()
		close(fileSizes)
	}()

	var numberOfFiles int32
	var totalFileSizes int64

loop:
	for {
		select {
		// test whether the fileSizes channel has been closed, using the two-result
		// form of receive operation.
		case fileSize, ok := <- fileSizes:
			if !ok {
				// Code arrive here when callingL close(fileSizes)
				// If the channel has been closed, the program breaks out of the loop. The labeled
				// break statement breaks out of both the select and the for loop; an unlabeled break
				// would break out of only the select, causing the loop to begin  the next iteration.

				// Print file totals. This is required because processing may complete and fileSizes
				// channel closes, but timer has yet some time to elapse
				printDiskUsage(numberOfFiles, totalFileSizes)
				break loop
			}

			// channel not closed yet, so update state
			numberOfFiles++
			totalFileSizes += fileSize
		case <- timer:
			printDiskUsage(numberOfFiles, totalFileSizes)
		}
	}
}

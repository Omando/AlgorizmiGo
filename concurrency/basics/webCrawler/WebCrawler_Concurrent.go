package webCrawler

import "fmt"

// Web crawling using breadth-first traversal
func CrawlConcurrently(topUrl string) {
	// This map is used to ensure that we do not crawl the same url twice
	processed := make(map[string]bool)

	// chUrls replaces the queue used in CrawlSequentially. Note that we use a
	// buffered channel to avoid a deadlock
	var chUrls chan []string = make(chan []string, 1)

	// Note use of a buffered channel to emulate a semaphore. The semaphore is used
	// to throttle i/o, without which we get errors such as too many open files
	var semaphore = make(chan struct{}, 10) // Maximum 10 concurrent requests

	// Populate the initial url
	// Without a buffered channel, the following code panics with this error:
	// fatal error: all goroutines are asleep - deadlock!
	// This is because a non-buffered channel blocks until another goroutine reads from it
	n := 1 // Initial number of url links
	chUrls <- []string{topUrl}

	for ; n > 0; n-- {
		urls := <-chUrls
		for _, url := range urls {
			if !processed[url] {
				n++
				fmt.Printf("Not seen %s. Current count %d\n", url, n)
				processed[url] = true

				// Crawl this url and return a list of new urls to crawl
				// The goroutine takes url as an explicit parameter to avoid the problem
				// of loop variable capture
				go func(u string) {
					// Use a semaphore to limit the number of open sockets
					semaphore <- struct{}{} // acquire token
					chUrls <- crawlUrl(url)
					<-semaphore // release token
				}(url)
			}
		}
	}

	fmt.Println("CrawlConcurrently has completed")
}

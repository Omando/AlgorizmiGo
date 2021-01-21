package webCrawler

import (
	"AlgorizmiGo/collections/sequential"
	"fmt"
)

// Web crawling using breadth-first traversal
func CrawlSequentially(topUrl string) {
	// This map is used to ensure that we do not crawl the same url twice
	processed := make(map[string]bool)

	// Used in breadth traversal
	var queue = sequential.CreateQueue()
	queue.Enqueue(topUrl)

	for !queue.IsEmpty() {
		item, err := queue.Dequeue()
		if err != nil {
			fmt.Printf("Failed to queue item: %v", err)
			continue
		}

		// Crawl this url and return a list of new urls to crawl
		url := item.(string)
		links := crawlUrl(url)
		processed[url] = true
		for _, link := range links {
			// Queue this link if not processed before
			if !processed[link] {
				queue.Enqueue(link)
			}
		}
	}
}

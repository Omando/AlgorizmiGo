package nonBlockingCache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func RunSequential(){
	var urls = []string {
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://godoc.org",
		"https://play.golang.org",
	}

	memo := New(HttpGetBody)
	for _, url := range urls {
		start := time.Now();
		value, err :=  memo.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s [%s]: %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func RunConcurrent(){
	var urls = []string {
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://godoc.org",
		"https://play.golang.org",
	}

	var wg sync.WaitGroup
	memo := New(HttpGetBody)
	for _, url := range urls {
		wg.Add(1)
		go func(link string) {
			start := time.Now();
			value, err := memo.Get(link)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s [%s]: %d bytes\n", link, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}

	// Wait until all urls have run
	wg.Wait()
}
package nonBlockingCache

import (
	"fmt"
	"log"
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

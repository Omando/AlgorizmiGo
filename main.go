package main

import (
	"AlgorizmiGo/concurrency/basics/chatServer"
	"flag"
	"fmt"
)

func main() {
	// To run: go run main.go -port 2222 -maxConn 2
	// Read port number from command line argument
	var port *int = flag.Int("port", 8000, "-port 1234")
	var maxConnections *int = flag.Int("maxConn", 2, "-maxConn 2")
	flag.Parse()
	fmt.Printf("Listening on port %d. Maximum number of connections =%d \n", *port, *maxConnections)

	//clockServer.Run(*port, *maxConnections)
	//echoServer.Run(*port, *maxConnections)
	//pipelines.Run()
	//pipelines.RunWithRange()
	//pipelines.RunConcurrentStages()
	//webCrawler.CrawlSequentially("https://golang.org")
	//webCrawler.CrawlConcurrently("https://golang.org")
	//counters.AbortableCountdown(10);
	//diretoryTraversal.TraverseRoots([]string{"C:\\Temp"})
	//diretoryTraversal.TraverseRootsAndMultiplexResults([]string{"C:\\Temp"})
	chatServer.Run(*port)

	fmt.Println("Completed")
}

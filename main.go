package main

import (
	"AlgorizmiGo/concurrency/basics/EchoServer"
	"flag"
	"fmt"
)

func main() {
	// To run: go run main.go -port 2222 -maxConn 2
	// Read port number from command line argument
	var port *int = flag.Int("port", 8000,"-port 1234")
	var maxConnections *int = flag.Int("maxConn", 2, "-maxConn 2")
	flag.Parse()
	fmt.Printf("Listening on port %d. Maximum number of connections =%d \n", *port, *maxConnections)

	//ClockServer.Run(*port, *maxConnections)
	EchoServer.Run(*port, *maxConnections)
}

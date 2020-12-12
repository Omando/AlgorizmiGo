package main

import (
	"AlgorizmiGo/concurrency/basics/ClockServer"
	"flag"
	"fmt"
)

func main() {
	// main -port 8000 -maxConn 2
	// Read port number from command line argument
	var port *int = flag.Int("port", 8000,"-port 1234")
	var maxConnections *int = flag.Int("maxConn", 2, "-maxConn 2")
	flag.Parse()
	fmt.Printf("Listening on port %d. Maximum number of connections =%d \n", *port, *maxConnections)

	ClockServer.Run(*port, *maxConnections)
}

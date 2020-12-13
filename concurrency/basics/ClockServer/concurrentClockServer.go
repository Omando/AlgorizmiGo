package ClockServer

import (
	"fmt"
	"log"
	"net"
)


func Run(port int, maxConnections int) {
	// Listen to incoming connections from clients
	address := fmt.Sprint("localhost:", port)
	listener, err := net.Listen("tcp", address)
	if (err != nil) {
		log.Fatal(err)		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	fmt.Printf("Listening on address: %s\n", address)

	// Accept upto maxConnection connections
	var connectionCount = 0
	for connectionCount <= maxConnections {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		connectionCount++
		go processConnection(connection)
	}
}

func processConnection(connection net.Conn) {
	// todo
}

package EchoServer

import (
	"fmt"
	"log"
	"net"
)

func Run(port int, maxConn int) {
	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on address %s\n", address)

	var connectionCount = 0
	for connectionCount <= maxConn {
		// listener.Accept blocks until an incoming connection request is made. A net.Conn object
		// representing the connection is returned
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received client connection%s\n", connection.LocalAddr().String())

		connectionCount++

		// Handle all incoming connection concurrently. Without 'go', only one client
		// is processed at a time since processConnection is a blocking operation that
		// returns only when the connection is terminated
		go processConnection(connection)
	}

}

func processConnection(connection net.Conn) {
	// todo
}

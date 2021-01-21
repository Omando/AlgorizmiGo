package clockServer

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// To run: go run main.go -port 2222 -maxConn 2
func Run(port int, maxConnections int) {
	// Listen to incoming connections from clients
	address := fmt.Sprint("localhost:", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err) // Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	fmt.Printf("Listening on address: %s\n", address)

	// Accept upto maxConnection connections
	var connectionCount = 0
	for connectionCount <= maxConnections {
		// listener.Accept blocks until an incoming connection request is made. A net.Conn object
		// representing the connection is returned
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		connectionCount++

		// Handle all incoming connection concurrently. Without 'go', only one client
		// is processed at a time since processConnection is a blocking operation that
		// returns only when the connection is terminated
		go processConnection(connection)
	}
}

// Handle a complete client connection. The infinite loop ends when a client disconnects
func processConnection(connection net.Conn) {
	defer connection.Close()

	// Write the current time to this connection every 1 second
	for {
		// Because net.Conn satisfies the io.Writer interface, the time is written directly to it
		// The time.Time.Format method provides a way to for mat date and time informatton by
		// example
		_, err := io.WriteString(connection, time.Now().Format("Mon Jan _2 15:04:05 2006\r"))
		if err != nil {
			fmt.Printf("Disconnecting from %s\n", connection.LocalAddr().String())
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

package ClockServer

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
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
	defer connection.Close()

	// Write the current time to this connection every 1 second
	for {
		_, err :=  io.WriteString(connection, time.Now().Format("Mon Jan _2 15:04:05 2006\r"))
		if err != nil {
			fmt.Printf("Disconnecting from %s\n", connection.LocalAddr().String())
			return		// client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

package chatServer

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Run(port int) {
	// Listen for incoming connections
	address := fmt.Sprintf("localhost:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	// Start broadcaster which maintains a set of connected clients and ensures
	// messages are broadcast to all connected clients
	go broadcaster()

	// Infinite loop to accept incoming client connections
	for {
		// Wait for and return the next connection
		conn, err := listener.Accept()
		if err != nil {
			// Log error then continue accepting connections
			log.Print(err)
			continue
		}

		// Handle all i/o for this client connection in a separate goroutine
		go handleConnection(conn)
	}
}

type client chan<-string         	// a string send-only channel
var entering = make(chan client)	// entering listens to messages of string send-only channel
var leaving = make(chan client)		// leaving listens to messages of string send-only channel
var message = make(chan string)		// message listens to messages of strings


/* broadcaster: listens on the global entering and leaving channels for announcements of arriving
and departing clients. When it receives one of these events, it updates the clients map, and if the
event was a departure, it closes the client’s send only channel. The broadcaster also listens for
events on the global messages channel, to which each client sends all its incoming messages.
When the broadcaster receives one of these events, it broadcasts the mess age to every connected client:
 */
func broadcaster() {
	// Create a map of send-only string channels to bool. This is the set
	// of connected clients
	clients:= make(map[client]bool)

	// infinitely multiplex select
	for {
		select {
		case msg := <- message:
			// Received a string message. Broadcast to all connected clients
			// Recall that each entry in clients is a send-only string channel
			// that was received using the entering channel
			for cli := range clients {
				cli <- msg
			}
		case clientChannel := <- entering:	// entering is a normal channel of 'send-only string channel'
			clients[clientChannel] = true
		case clientChannel := <- leaving:	// leaving is a normal channel of 'send-only string channel'
			delete(clients, clientChannel)
			close(clientChannel)
		}
	}
}

/* The handleConn function creates a new outgoing message channel for its client and announces
the arrival of this client to the broadcaster over the entering channel. Then it reads every line
of text from the client, sending each line to the broadcaster over the global incoming message
channel, prefixing each message with the identity of its sender. Once there is nothing more to
read from the client, handleConn announces the departure of the client over the leaving channel
and closes the connection */
func handleConnection(conn net.Conn) {
	// Create a string channel. This channel will be sent via
	// the global entering channel so that the broadcaster can cache it
	// and later use it to broadcast messages back to this client
	var outgoing = make(chan string)

	// goroutine for each message broadcast to the client’s out going message channel
	// When the broadcaster broadcasts to all connected clients, it writes to this client'
	// outgoing message. The goroutine below capture this information and displays in
	// the client's network connection
	// The client writer’s loop terminates when the broadcaster closes the channel after
	// receiving a leaving notification from this client
	go func() {
		for msg := range outgoing {
			fmt.Fprintf(conn, "%s\r\n", msg)
		}
	}()

	whoAmI := conn.RemoteAddr().String()
	outgoing <- "You are: " + whoAmI	// client sends message to itself identifying its id
	message <- whoAmI + " has arrived"	// Let all clients know (via the broadcaster) that a new client is available
	entering <- outgoing				// announce this client to the broadcaster

	// Read every line  of text from the client, sending each line to the broadcaster over the
	//global incoming message channel, prefixing each message with the identity of its sender
	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- whoAmI + ": " + input.Text()
	}

	// Client is done sending messages. Ask broadcaster to remove this client
	// The broadcaster will close this client's channel and the range call
	// in the goroutine above will terminated
	leaving <- outgoing;
	message <- whoAmI + " has left"
	conn.Close()
}

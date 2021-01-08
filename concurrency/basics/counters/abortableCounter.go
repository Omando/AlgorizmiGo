package counters

import (
	"fmt"
	"os"
	"time"
)

func AbortableCountdown(countdownFrom int) {
	var abort chan struct{} = make(chan struct{})

	// Wait for the user to enter a character to abort the countdown
	go func() {
		// Read a single byte from the command line console
		os.Stdin.Read(make([]byte, 1))

		// Send a signal on the abort channel
		abort <- struct{}{}
	}()

	// Start the timed countdown and monitor for abort signals
	var tick *time.Ticker = time.NewTicker(1 * time.Second)
	for i := countdownFrom; i >= 0; i-- {
		select {
		case <- tick.C:
			fmt.Println(i)
		case <- abort:
			fmt.Println("Countdown aborted")
			return
		}
	}

	// Turn off the ticket
	tick.Stop()
}

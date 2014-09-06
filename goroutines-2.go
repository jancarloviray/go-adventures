/*
CHANNELS
	- a channel is a Go language construct that provides
	a mechanism for two goroutines to synchronize execution
	and communicate by passing a value of a specified
	element type.

	- without channels, go will not block meaning go will
	not wait until goroutines finish. the function will
	just execute and then exits (unless a timer is set),
	leaving the goroutine in another thread "twilight zone"

	- The <- operator specifies the channel direction,
	send or receive. If no direction is given, the channel
	is bi-directional

	`chan T` - can be used to send and receive values
	of type T
	`chan<-float64` - can be used to send float64s
	`<-chan int` - can be used to receive ints
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)

	go func() {
		// simulate long running process
		time.Sleep(2000 * time.Millisecond)

		// simulate results
		fmt.Println("goroutine has been run")
		c <- true
	}()

	fmt.Println("program has exited")

	// by having this line, go waits or "blocks" until
	// the channel has a value. taking this line out
	// will not print "goroutine has been run"
	<-c
}

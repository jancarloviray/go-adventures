/*
Buffered Channel

	- by default, go channels can only contain one value. It will
	block until it frees that value, to which will allow itself
	to contain another value

	- in other words, channels by default can only hold 8oz of
	water. if you try putting more to it, you cannot because
	it's already full. you can only put additional water if you
	take out the existing one. if you want it to be able to
	contain more, then use "buffered channel"

*/

package main

import (
	"fmt"
)

func main() {
	// unbuffered channel
	c := make(chan bool)

	// buffered channel
	//c := make(chan bool, 3)

	go func() {
		// simulate results
		c <- true

		// execution will not reach here
		// since <-c will not block the
		// execution to wait for the other
		// two c<-
		c <- true
		c <- true
		fmt.Println("goroutine has been run")
	}()

	fmt.Println("program has exited")

	// note that "c" is not a buffered channel,
	// you will not see "goroutine has been run.."
	// printed since it assumes by default to only
	// contain one
	fmt.Println(<-c)
}

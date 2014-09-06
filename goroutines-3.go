/*
CHANNELS
	- a channel is a Go language construct that provides
	a mechanism for two goroutines to synchronize execution
	and communicate by passing a value of a specified
	element type.
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
)

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		//note: this will not work
		//due to closure situation (similar to js)

		//go func(){
		//	c <- i
		//}()

		go func(i int) {
			c <- i
		}(i)
	}

	fmt.Println(<-c, <-c, <-c, <-c) // 0, 1, 2, 3
}

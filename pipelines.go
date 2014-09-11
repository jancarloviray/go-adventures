/*
Pipeline
	- a series of stages connected by channels where each stage is a group
	of goroutines running the same function
*/

package main

import (
	"fmt"
)

func main() {
	// set up pipeline
	c := gen(2, 3)
	out := sq(c)

	// consume
	fmt.Println(<-out) // 4 (2*2)
	fmt.Println(<-out) // 9 (3*3)
}

// first stage
// a function that converts a list of integers to a channel that emits
// integers in the list.
// it starts a goroutine that sends integers on the channel and closes
// the channel when all the values have been sent

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// second stage
// receives integers from a channel and returns a channel that emits
// the square of each received integer.
// after the inbound channel is closed and this stage has sent all the
// values downstream, it closes the outbound channel.

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

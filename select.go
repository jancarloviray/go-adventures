package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c3 <- "three"
	}()

	for {
		select {
		case m1 := <-c1:
			fmt.Println("received ", m1)
		case m2 := <-c2:
			fmt.Println("received ", m2)
		case m3 := <-c3:
			fmt.Println("received ", m3)
		}
	}
}

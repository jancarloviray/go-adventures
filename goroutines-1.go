/*
Placing a "go" statement before a function call starts
the execution of that function as an independent
concurrent thread in the same address space as the
calling code.

Such thread is called "goroutine" in Golang.

Goroutines are lightweight, costing little more
than the allocation of stack space. The stacks start
small and grow by allocating and freeing heap storage
as required.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	//Lesson1()
	//Lesson2()
	Lesson3("I am Lesson 3!", 2*time.Second)

	// No output? Uncomment this. Executing
	// functions do not wait for goroutines
	// to finish.
	// time.Sleep(3 * time.Second)
}

// Lesson1 showcases basic goroutine.
// Note that "Hello there!" might or might not print.
// When you start a goroutine, the calling code does
// not wait for the goroutine to finish, but continues
// calling further.
// To synchronize execution, you must use "Channels"
func Lesson1() {
	go fmt.Println("Hello there, from Lesson 1!")
	fmt.Println("Hello world!")
}

// Lesson2 further reinforces that the goroutines
// do execute in the background.
// Because we have a timer, it is sure that we will
// see the "Hello there, from Lesson 2" printed in
// the console.
func Lesson2() {
	go fmt.Println("Hello there, from Lesson 2!")
	fmt.Println("Hello world!")

	// wait for 1 second for other goroutine to finish
	time.Sleep(time.Second)
}

func Lesson3(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println("NOTICE:", text)
	}()
}

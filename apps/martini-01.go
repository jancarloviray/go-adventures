package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()

	// A Handler in Martini is any callable function.
	// If your function returns something, Martini will
	// write it out to the HTTP response body.

	m.Get("/", func() string {
		return "Hello World!"
	})

	// In addition, a handler function can return
	// a (int, string) and Martini will write out a response
	// code as well as a body
	m.Get("/error", func() (int, string) {
		return 500, "Error in Server"
	})

	m.Run()
}

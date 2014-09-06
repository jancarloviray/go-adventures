/*
Errors

By default, errors are the last return value
and have type `error`, a built-in interface

errors.New constructs a basic `error` value with
the given error message

A `nil` value in the error position indicates that
there was no error

It's possible to use custom types as `errors` by
implementing Error() method on them.
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(err2(42))
}

func err1(arg int) (int, error) {
	if arg == 42 {
		return arg, errors.New("can't work with 42")
	}
	return arg + 10, nil
}

/*
Custom types implemented as error by implementing
Error() method on them.
*/

type argumentError struct {
	arg  int
	prob string
}

func (e *argumentError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func err2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argumentError{arg, "can't work with 42"}
	}
	return arg + 3, nil
}

package main

import (
	"fmt"
)

func main() {
	animals := []Animal{Dog{}, Cat{}, JavaProgrammer{}}
	for k, v := range animals {
		fmt.Println(k, v.Speak())
	}
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

type JavaProgrammer struct{}

func (j JavaProgrammer) Speak() string { return "Design patterns!" }

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		strings.Contains("test", "es"),                // true
		strings.Count("test", "t"),                    // 2
		strings.EqualFold("Go", "go"),                 // true
		strings.Fields("foo bar  baz hi"),             // [foo bar baz hi]
		strings.HasPrefix("test", "te"),               // true
		strings.HasSuffix("test", "st"),               // true
		strings.Index("chicken", "ken"),               // 4
		strings.Join([]string{"hello", "world"}, "|"), // hello|world
		strings.LastIndex("go gopher", "go"),          // 3
		"more at http://golang.org/pkg/strings/",
	)
}

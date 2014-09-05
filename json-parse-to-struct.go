// PARSE JSON TO STRUCT

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s ServerList
	/*
		NOTE: backticks are alternative to string syntax
		and are essentially the same as "jons:\"type\""
	*/
	str := `{"servers":[{"serverName":"Test1","someKey":"hello!","serverIP":"127.0.0.1"},{"serverName":"Test2","serverIP":"127.0.0.2"}]}`

	// func Unmarshal(data []byte, v interface{}) error
	json.Unmarshal([]byte(str), &s)

	fmt.Println(s)
}

type ServerList struct {
	Servers []Server
}

/*
NOTE: Go only assigns fields that can be found
in the json string and ignores all others. This
essentially is the same as _.pick() in underscore.
*/
type Server struct {
	ServerName string
	ServerIP   string
}

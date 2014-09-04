package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func main(){
	// set router
	http.HandleFunc("/", sayHello) 
	// set listen port
	err := http.ListenAndServe(":9090", nil) 
	if err != nil { log.Fatal("ListenAndServe: ", err) }
}

func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm() // parse arguments
	fmt.Println(r.Form) // print form info
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello!") // send to client
}
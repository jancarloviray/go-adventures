package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	dst := "http://www.golang.org"
	src := ":80"

	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		dst = args[0]
	}

	if len(args) == 2 {
		src = args[1]
	}

	u, e := url.Parse(dst)

	if e != nil {
		log.Fatal("Bad destination.")
	}
	// func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy
	//     NewSingleHostReverseProxy returns a new ReverseProxy that
	//     rewrites URLs to the scheme, host, and base path provided
	//     in target. If the target's path is "/base" and the incoming
	//     request was for "/dir", the target request will be for
	//      /base/dir.
	h := httputil.NewSingleHostReverseProxy(u)

	s := &http.Server{
		Addr:    src,
		Handler: h,
		//MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

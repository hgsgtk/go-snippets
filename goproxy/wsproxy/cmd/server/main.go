package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hgsgtk/go-snippets/goproxy/wsproxy"
)

func main() {
	p := 12345
	log.Printf("Starting websocket echo server on port %d", p)

	http.HandleFunc("/", wsproxy.Echo)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), nil); err != nil {
		log.Panicf("Error while starting to listen: %#v", err)
	}
}

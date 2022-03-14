package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Printf("error on listening a random port: %s\n", err)
		return
	}

	port := l.Addr().(*net.TCPAddr).Port
	log.Printf("using port: %d", port)
	if err := http.Serve(l, nil); err != nil {
		log.Printf("error on serving: %s\n", err)
		return
	}
}

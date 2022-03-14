package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Printf("error on listening a random port: %s\n", err)
		os.Exit(1)
	}

	log.Printf("using port: %d", l.Addr().(*net.TCPAddr).Port)
	if err := http.Serve(l, nil); err != nil {
		log.Printf("error on serving: %s\n", err)
		os.Exit(1)
	}
}

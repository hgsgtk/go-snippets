package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

// https://pkg.go.dev/golang.org/x/net/websocket#example-Dial
func main() {
	origin := "http://localhost/"
	url := "ws://localhost:12345/echo"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	ws.Write([]byte("hello, world!\n"))

	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s\n", msg[:n])
}

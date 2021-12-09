package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrading error: %#v\n", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Reading error: %#v\n", err)
			break
		}
		log.Printf("recv: message %q", message)
		if err := c.WriteMessage(mt, message); err != nil {
			log.Printf("Writing error: %#v\n", err)
			break
		}
	}
}

func main() {
	p := 12345
	log.Printf("Starting websocket echo server on port %d", p)

	http.HandleFunc("/", echo)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), nil); err != nil {
		log.Panicf("Error while starting to listen: %#v", err)
	}
}

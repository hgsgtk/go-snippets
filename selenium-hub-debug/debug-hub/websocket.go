package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func StartWebSocketClient(host, port string) {
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", host, port), Path: "/"}
	log.Printf("Connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	log.Println("Connected to the EventBus")

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("Received: %s", message)
	}
}

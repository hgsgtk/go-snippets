package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	endpointUrl := "ws://localhost:12345"

	dialer := websocket.Dialer{}
	// TODO: What contains in a response
	c, _, err := dialer.Dial(endpointUrl, nil)
	if err != nil {
		log.Panicf("Dial failed: %#v\n", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer c.Close()
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("read message: %#v\n", err)
				return
			}
			log.Printf("recv: %s\n", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			if err := c.WriteMessage(websocket.TextMessage, []byte(t.String())); err != nil {
				log.Printf("writing: %#v\n", err)
				return
			}
		case <-interrupt:
			log.Println("interrupting")
			if err := c.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(
					websocket.CloseNormalClosure, "",
				)); err != nil {
				log.Printf("error closing: %#v", err)
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}
}

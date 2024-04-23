package main

import (
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, err := zmq.NewContext()
	if err != nil {
		log.Fatalf("Error creating new context: %s\n", err)
	}

	s, err := zctx.NewSocket(zmq.REP)
	if err != nil {
		log.Fatalf("Error creating new socket: %s\n", err)
	}
	addr := "tcp://*:5556"
	log.Printf("Start zeromq server %s\n", addr)
	if err := s.Bind(addr); err != nil {
		log.Fatalf("Error binding to the server: %s\n", err)
	}
	log.Println("Server started...")

	for {
		// Wait for next request from client
		msg, err := s.Recv(0)
		if err != nil {
			log.Fatalf("Error receiving request: %s\n", err)
		}
		log.Printf("Received %s\n", msg)

		// Do some 'work'
		time.Sleep(time.Second * 1)

		// Send reply back to client
		if _, err := s.Send("World", 0); err != nil {
			log.Fatalf("Error sending reply: %s\n", err)
		}
	}
}

package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)
func StartZeroMQClient(host, port string) {
  zctx, err := zmq.NewContext()

  if err != nil {
    fmt.Printf("Error creating new context: %s\n", err)
    return
  }

  // Socket to talk to server
  addr := fmt.Sprintf("tcp://%s:%s", host, port)
  fmt.Printf("Connecting to %s...\n", addr)
  s, err := zctx.NewSocket(zmq.REQ)
  if err != nil {
    fmt.Printf("Error creating new socket: %s\n", err)
    return
  }
  if err := s.Connect(addr); err != nil {
    fmt.Printf("Error connecting to the server: %s\n", err)
    return
  }

  // Do 10 requests, waiting each time for a response
  for i := 0; i < 10; i++ {
          fmt.Printf("Sending request to %s %d...\n", addr, i)
          if _, err := s.Send("Hello", 0); err != nil {
            fmt.Printf("Error sending request: %s\n", err)
            break
          }

          msg, err := s.Recv(0)
          if err != nil {
            fmt.Printf("Error receiving reply: %s\n", err)
            break
          }
          fmt.Printf("Received reply from %s  %d [ %s ]\n", addr, i, msg)
  }
}

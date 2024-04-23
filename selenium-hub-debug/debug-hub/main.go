package main

import (
	"os"
	"time"
)

func main() {
	go startServer("4444")
  go startServer("5553")

  if os.Getenv("ENABLE_EVENT_BUS") == "true" {
    host := os.Getenv("SE_EVENT_BUS_HOST")
    publishPort := os.Getenv("SE_EVENT_BUS_PUBLISH_PORT")
    busSubscribePort := os.Getenv("SE_EVENT_BUS_SUBSCRIBE_PORT")

    time.Sleep(5 * time.Second) // Wait for EventBus to start

    go StartZeroMQClient(host, busSubscribePort)
    go StartZeroMQClient(host, publishPort)
  }
	select {}
}

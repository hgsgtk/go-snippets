package main

import (
	"log"
)

type Message struct {
	Body string
}

func main() {
	ch := make(chan Message)

	go func() {
		ch <- Message{
			Body: "hoge",
		}
	}()

	m := <-ch
	log.Printf("message %q received", m)

	go func() {
		ch <- Message{
			Body: "huga",
		}
	}()
	m = <-ch
	log.Printf("message %q received", m)
	// 2022/02/05 13:47:53 message {"hoge"} received
	// 2022/02/05 13:47:53 message {"huga"} received
}

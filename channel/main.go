package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Message struct {
	Body string
}

func main() {
	ch := make(chan Message)

	go func() {
		defer fmt.Fprintln(os.Stdout, "message hoge sent")

		time.Sleep(1 * time.Second)
		ch <- Message{
			Body: "hoge",
		}
	}()

	go func() {
		defer fmt.Fprintln(os.Stdout, "message huga sent")

		time.Sleep(2 * time.Second)
		ch <- Message{
			Body: "huga",
		}
	}()

	time.Sleep(3 * time.Second)

	m := <-ch
	log.Printf("message %q received", m)
	m = <-ch
	log.Printf("message %q received", m)

	time.Sleep(1 * time.Second)

	/*
		goroutine 2 is blocked until main goroutine receives the value.
		$ go run main.go
		message hoge sent
		2022/02/05 13:55:35 message {"hoge"} received
		2022/02/05 13:55:35 message {"huga"} received
		message huga sent
	*/
}

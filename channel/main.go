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
	ch := make(chan Message, 2) // buffered 2

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

	go func() {
		defer fmt.Fprintln(os.Stdout, "message piyo sent")

		time.Sleep(2 * time.Second)
		ch <- Message{
			Body: "piyo",
		}
	}()

	time.Sleep(3 * time.Second)

	for {
		m := <-ch
		log.Printf("message %q received", m)
	}

	/*
		buffered channel
		$ go run main.go
		message hoge sent
		message huga sent
		message piyo sent
		2022/02/05 13:58:40 message {"hoge"} received
		2022/02/05 13:58:40 message {"huga"} received
		2022/02/05 13:58:40 message {"piyo"} received
		fatal error: all goroutines are asleep - deadlock!
	*/
}

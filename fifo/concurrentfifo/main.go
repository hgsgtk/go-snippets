package main

import (
	"fmt"
	"time"

	"github.com/enriquebris/goconcurrentqueue"
)

type Struct struct {
	A string
	B int
}

func main() {
	// unfixedFIFO()

	waitForDequeue()
}

func waitForDequeue() {
	fifoQ := goconcurrentqueue.NewFIFO()
	done := make(chan struct{})

	go func() {
		fmt.Println("waiting for next element...")
		item, err := fifoQ.DequeueOrWaitForNextElement()
		if err != nil {
			panic(err)
		}
		fmt.Println("got next element:", item)

		done <- struct{}{}
	}()

	fmt.Println("go to sleep for a while...")
	time.Sleep(3 * time.Second)

	fmt.Println("Enqueueing new element...")
	fifoQ.Enqueue("new element")

	<-done
}

func unfixedFIFO() {
	q := goconcurrentqueue.NewFIFO()

	q.Enqueue("any string value")
	q.Enqueue(5)
	q.Enqueue(Struct{A: "any string value", B: 5})

	fmt.Printf("queue size: %d\n", q.GetLen())

	item, err := q.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dequeued item: %v\n", item)

	fmt.Printf("queue size: %d\n", q.GetLen())

	item, err = q.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dequeued item: %v\n", item)

	item, err = q.Dequeue()
	if err != nil {
		panic(err)
	}
	fmt.Printf("dequeued item: %v\n", item)
}

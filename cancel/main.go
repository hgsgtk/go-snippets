package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func execute(ctx context.Context) error {
	proc1 := make(chan struct{}, 1)
	proc2 := make(chan struct{}, 1)

	go func() {
		// Would be done before timeout
		time.Sleep(1 * time.Second)
		proc1 <- struct{}{}
	}()

	go func() {
		// Would not be executed because timeout comes first
		time.Sleep(3 * time.Second)
		proc2 <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-proc1:
			fmt.Println("process 1 done")
		case <-proc2:
			fmt.Println("process 2 done")

		}
	}

	return nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := execute(ctx); err != nil {
		log.Fatalf("error: %#v\n", err)
	}
	log.Println("Success to process in time")
}

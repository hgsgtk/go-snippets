package main

import (
	"fmt"
	"sync/atomic"
)

var counter1 uint64
var counter2 uint64
var counter3 uint64

func procChannel() {
	chs := [3]chan struct{}{}
	for i := 0; i < 3; i++ {
		chs[i] = make(chan struct{}, 1)
		chs[i] <- struct{}{}
	}

	// Already sent to three channels
	select {
	case <-chs[0]:
		atomic.AddUint64(&counter1, 1)
	case <-chs[1]:
		atomic.AddUint64(&counter2, 1)
	case <-chs[2]:
		atomic.AddUint64(&counter3, 1)
	}
}

func main() {
	for i := 0; i < 500000; i++ {
		procChannel()
	}

	fmt.Printf(
		"1: %d times, 2: %d times, 3: %d times\n",
		counter1, counter2, counter3,
	)
}

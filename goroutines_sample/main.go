package main

import (
	"fmt"
	"runtime"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println("num of goroutine: ", runtime.NumGoroutine())
	select {}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

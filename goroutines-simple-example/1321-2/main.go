package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		// goroutineの起動はOSのネイティブスレッドより高速だけど、コストゼロというわけではない
		go func() {
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}

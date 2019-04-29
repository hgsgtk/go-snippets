package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//func main() {
//	defer profile.Start(profile.ProfilePath(".")).Stop()
//
//	fmt.Println("start")
//	for i := 0; i < 1000; i++ {
//		fib(30)
//	}
//}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("start")
	for {
		fib(30)
	}
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.WaitGroup
	// 多数のgoroutineで実行しているジョブの終了待ちに使う
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("work 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("work 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("finish")
}

package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}

func main() {
	// チャネルが有用な用途: データの所有権を渡す場合、作業を並列化して分散する場合、非同期で結果を受け取る場合
	// Mutexが有用な用途: キャッシュ、状態管理
	// See also https://github.com/golang/go/wiki/MutexOrChannel
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
}

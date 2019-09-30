package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 依存関係のあるタスクを表現する

func readFile(path string) chan string {
	promise := make(chan string)
	go func() {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("read error %s\n", err.Error())
			close(promise)
		} else {
			promise <- string(content)
		}
	}()
	return promise
}

func printFunc(futureSource chan string) chan []string {
	promise := make(chan []string)
	go func() {
		var result []string

		for _, line := range strings.Split(<-futureSource, "\n") {
			if strings.HasPrefix(line, "func ") {
				result = append(result, line)
			}
		}
		promise <- result
	}()
	return promise
}

func main() {
	// ファイルの読み込みが終わった時点でそれが可能されるFutureを返す
	futureSource := readFile("future_promise.go")
	// ソースを受け取り関数宣言のリストを返す
	futureFuncs := printFunc(futureSource)
	fmt.Println(strings.Join(<-futureFuncs, "\n"))
}

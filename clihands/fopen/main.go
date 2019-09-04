package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("hoge.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ft := bufio.NewScanner(f)
	for ft.Scan() {
		fmt.Println(ft.Text())
	}
}

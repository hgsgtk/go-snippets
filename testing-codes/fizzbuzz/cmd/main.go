package main

import (
	"fmt"
	"github.com/hgsgtk/go-snippets/fizzbuzz"
	"os"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Fprintf(
			os.Stdout,
			"number: %d, result: %s\n",
			i, fizzbuzz.Run(i))
	}
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Printf("Please input a text which you want to count. \n")
	fmt.Scan(&str)
	fmt.Printf("Length: %d \n", utf8.RuneCountInString(str))
}

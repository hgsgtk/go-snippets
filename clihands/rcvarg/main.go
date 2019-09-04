package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	// go run main.go
	// Output: [/var/folders/lg/rdr0tvnd6kzblb0y1xmpvvx00000gn/T/go-build546768911/b001/exe/main]
	// １つ目の要素はプログラム名
}

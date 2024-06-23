package main

import (
	"fmt"

	"github.com/nxadm/tail"
)

func main() {
	// Create a tail
	t, err := tail.TailFile(
		"./selenium.log", tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Print the text of each received line
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "Hello")
	output, err := cmd.Output()
	if err != nil {
		panic(err) // 稼働中の本番サービスなどではpanicしない
	}
	fmt.Print(string(output))
}

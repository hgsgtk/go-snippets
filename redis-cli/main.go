package main

import (
	"fmt"
	"os"

	"github.com/hgsgtk/go-snippets/redis-cli/persistence/kvs"
)

func main() {
	client, err := kvs.New("localhost:6379")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	defer client.Close()

	if err := kvs.SetToken(client, "key1", 1); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	id, err := kvs.GetIDByToken(client, "key1")
	if err == kvs.Nil {
		fmt.Fprintln(os.Stderr, "no value")
		return
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	fmt.Fprintln(os.Stdout, id)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

// https://christina04.hatenablog.com/entry/golang-pprof-memory
// pprof -http=:9000 http://localhost:6060/debug/pprof/heap
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("start")
	leak()
	fmt.Println("end")

}

func leak() {
	var global = make(map[string]string)

	addMap(global)
}

func addMap(s map[string]string) {
	i := int64(0)
	for {
		key := "key" + strconv.FormatInt(i, 10)
		val := "value" + strconv.FormatInt(i, 10)
		s[key] = val
		i++
	}
}

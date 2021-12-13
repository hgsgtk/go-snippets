package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	port := 54321
	log.Printf("Starting proxy server on port %d", port)

	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.Verbose = true

	err := http.ListenAndServe(":54321", proxy)
	if err != nil {
		log.Fatal(err.Error())
	}
}

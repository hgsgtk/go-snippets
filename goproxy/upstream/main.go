package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/elazarl/goproxy.v1"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	os.Setenv("HTTP_PROXY", "http://localhost:32730")

	log.Fatal(http.ListenAndServe(":30000", proxy))	
}

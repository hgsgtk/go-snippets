package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/elazarl/goproxy"
)

func main() {
	url, err := url.Parse("http://username:password@localhost:32730")
	if err != nil {
		panic(err)
	}
	os.Setenv("HTTP_PROXY", "http://username:password@localhost:32730")
	os.Setenv("HTTPS_PROXY", "http://username:password@localhost:32730")

	log.Println(url.Scheme)
	log.Println(url.Host)
	log.Println(url.User.Username())
	log.Println(url.User.Password())

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.KeepHeader = true

	log.Fatal(http.ListenAndServe(":30000", proxy))	
}

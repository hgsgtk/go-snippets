package main

import (
	"log"
	"net/http"

	"gopkg.in/elazarl/goproxy.v1"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	/*
		AlwaysReject is a HttpsHandler that drops any CONNECT request

		$ curl -I -Lv -x http://127.0.0.1:8080 \
		https://example.com

		2022/01/18 15:59:32 [001] INFO: Running 1 CONNECT handlers
		2022/01/18 15:59:32 [001] INFO: on 0th handler: &{1 <nil> 0x1022c7c60} example.com:443
	*/
	proxy.OnRequest().HandleConnect(goproxy.AlwaysReject)

	log.Fatal(http.ListenAndServe(":8080", proxy))
}

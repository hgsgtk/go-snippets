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
		 AlwaysMitm is a HttpsHandler that always eavesdrop https connections

		$ curl -I -Lv -x http://127.0.0.1:8080 \
		https://example.com

		2022/01/18 15:47:00 [001] WARN: Cannot handshake client example.com:443 remote error: tls: unknown certificate authority
	*/
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	log.Fatal(http.ListenAndServe(":8080", proxy))
}

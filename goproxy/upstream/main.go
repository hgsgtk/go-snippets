package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/elazarl/goproxy"
)

const (
	ProxyAuthHeader = "Proxy-Authorization"
)

func SetBasicAuth(username, password string, req *http.Request) {
	req.Header.Set(ProxyAuthHeader, fmt.Sprintf("Basic %s", basicAuth(username, password)))
}

func basicAuth(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

const (
	username = "username"
	password = "password"
	proxyURL = "http://localhost:32730"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.Tr.Proxy = func(req *http.Request) (*url.URL, error) {
		return url.Parse(proxyURL)
	}
	connectReqHandler := func(req *http.Request) {
		SetBasicAuth(username, password, req)
	}
	proxy.ConnectDial = proxy.NewConnectDialToProxyWithHandler(proxyURL, connectReqHandler)
	proxy.OnRequest().Do(goproxy.FuncReqHandler(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		SetBasicAuth(username, password, req)
		return req, nil
	}))

	log.Fatal(http.ListenAndServe(":30000", proxy))	
}

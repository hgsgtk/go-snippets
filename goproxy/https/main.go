// https://github.com/elazarl/goproxy/tree/d06c3be7c11b750d8cd76d0f094936e07cac0ada/examples/goproxy-transparent
// TODO: support WCCP means...?
// TODO: explicit means...? see more https://github.com/elazarl/goproxy/blob/d06c3be7c11b750d8cd76d0f094936e07cac0ada/examples/goproxy-eavesdropper/main.go .
// TODO: SNI value in the TLS ClientHello which most modern clients do these days
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"

	"github.com/elazarl/goproxy"
)

func main() {
	httpAddr := ":3128"
	httpsAddr := ":3129"

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Printf(
		"Server staring up! - configured to listen on HTTP interface %s and HTTPS interface %s",
		httpAddr,
		httpsAddr,
	)

	// TODO: NonproxyHandler means...
	proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Host == "" {
			fmt.Fprintln(w, "Cannot handle requests without Host header, e.g., HTTP 1.0")
			return
		}
		r.URL.Scheme = "http"
		r.URL.Host = r.Host
		// TODO: how can we use it?
		// Standard net/http function. Shouldn't be used directly, http.Serve will use it.
		proxy.ServeHTTP(w, r)
	})

	// TODO: ^.*$ means? => all hosts? is it needed?
	//  AlwaysMitm is a HttpsHandler that always eavesdrop https connections
	proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*$"))).
		HandleConnect(goproxy.AlwaysMitm)
	// TODO: HTTP can be recognized by only port numbers?
	proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*:80$"))).
		// TODO: HijackConnect
		// TODO: client is for what?
		HijackConnect(func(r *http.Request, client net.Conn, ctx *goproxy.ProxyCtx) {
			defer func() {
				if e := recover(); e != nil {
					ctx.Logf("error connection to remote: %v", e)
					client.Write([]byte("HTTP/1.1 500 Connot reach destination\r\n\r\n"))
				}
			}()

			// send the request to remote and receive response from remote
			// TODO: client can act reader and writer?
			clientBuf := bufio.NewReadWriter(
				bufio.NewReader(client),
				bufio.NewWriter(client),
			)
			remote, err := connectDial(proxy, "tcp", r.URL.Host)
			if err != nil {
				// TODO: better error handling is?
				panic(err)
			}
			// remote is reader and writer also. The reason is...?
			remoteBuf := bufio.NewReadWriter(
				bufio.NewReader(remote),
				bufio.NewWriter(remote),
			)

			for {
				// TODO: proper error handling...?
				// HTTP wire text is contained in the buffer.
				// ReadRequest reads and parses an incoming request from b.
				r, err := http.ReadRequest(clientBuf.Reader)
				if err != nil {
					panic(err)
				}
				// TODO: what is it?
				if err := r.Write(remoteBuf); err != nil {
					panic(err)
				}
				if err := remoteBuf.Flush(); err != nil {
					panic(err)
				}
				resp, err := http.ReadResponse(remoteBuf.Reader, r)
				if err != nil {
					panic(err)
				}
				// TODO: what's is it?
				if err := resp.Write(clientBuf.Writer); err != nil {
					panic(err)
				}
				if err := clientBuf.Flush(); err != nil {
					panic(err)
				}

			}
		})

	go func() {
		log.Fatalln(http.ListenAndServe(httpAddr, proxy))
	}()

	// TODO: listen to the TLS ClientHello
}

func dial(proxy *goproxy.ProxyHttpServer, network, addr string) (c net.Conn, err error) {
	// Tr: HTTP Transport
	// TODO: Understand HTTP transport deeply
	if proxy.Tr.Dial != nil {
		// TODO: use DialContext instead
		// TODO: make a pull request to use a context
		return proxy.Tr.Dial(network, addr)
	}
	// TODO: What's difference between proxy.Tr.Dial and net.Dial
	return net.Dial(network, addr)
}

// TODO: what's it?
func connectDial(proxy *goproxy.ProxyHttpServer, network, addr string) (c net.Conn, err error) {
	// ConnectDial will be used to create TCP connections for CONNECT requests if nil Tr.Dial will be used
	if proxy.ConnectDial == nil {
		return dial(proxy, network, addr)
	}
	return proxy.ConnectDial(network, addr)
}

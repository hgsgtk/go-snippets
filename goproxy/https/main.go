// Reference sample codes
// - https://github.com/elazarl/goproxy/blob/adb46da277acd7aea06aeb8b5a21ec6bef7fb247/examples/goproxy-transparent/transparent.go
// - https://www.agwa.name/blog/post/writing_an_sni_proxy_in_go/media/sniproxy.go
// TODO: support WCCP means...?
// TODO: explicit means...? see more https://github.com/elazarl/goproxy/blob/d06c3be7c11b750d8cd76d0f094936e07cac0ada/examples/goproxy-eavesdropper/main.go .
// TODO: SNI value in the TLS ClientHello which most modern clients do these days
package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"

	"github.com/elazarl/goproxy"
	// TODO: replace go-vhost to the crypto/tls package
	"github.com/inconshreveable/go-vhost"
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
			remote, err := connectDial(r.Context(), proxy, "tcp", r.URL.Host)
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

	// A standard listen/accept loop to create a TCP server https://pkg.go.dev/net#pkg-overview
	//
	// listen to the TLS ClientHello
	// TODO: should support non-SNI request? https://github.com/elazarl/goproxy/issues/231
	ln, err := net.Listen("tcp", httpsAddr)
	if err != nil {
		log.Fatalf("Error listening for https connection - %v", err)
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting new connection - %v", err)
			continue
		}
		// Why goroutine?
		// Note: once the Client Hello is read, the bytes are gone
		// 	quoted from https://www.agwa.name/blog/post/writing_an_sni_proxy_in_go
		go func(c net.Conn) {
			// TLS parses the ClientHello message on conn and returns a new.
			// Client Hello is defined at RFC 8446
			// https://datatracker.ietf.org/doc/html/rfc8446#section-4.1.2
			// > When a client first connects to a server, it is REQUIRED to send the
			// > ClientHello as its first TLS message.
			tlsConn, err := vhost.TLS(c)
			if err != nil {
				log.Printf("Error accepting new connection - %v", err)
				// TODO: No need to continue?
			}
			// Host() returns the server name only set if the client is using SNI.
			if tlsConn.Host() == "" {
				// TODO: non-SNI enabled clients...?
				//
				// # SNI (Server Name Indication)
				// In Client Hello, the client greets the server and tells it the server name it wants to connect.
				// This is called is SNI (Server Name Indication) which is defined at RFC6066.
				// https://datatracker.ietf.org/doc/html/rfc6066#section-3 (Transport Layer Security (TLS) Extensions: Extension Definitions)
				// > TLS does not provide a mechanism for a client to tel a server the name of the server it is contacting.
				//
				// # Use cURL with SNI (Server name indication)
				// https://stackoverflow.com/questions/12941703/use-curl-with-sni-server-name-indication
				// It didn't work to add the option `--resolve`.
				// i.e.
				// 		$ curl -vik \--resolve localhost:3129:127.0.0.1:3129 \
				//		-x http://localhost:3129 \
				//		https://example.com
				//
				// 		2022/01/17 14:43:49 Connot support non-SNI enabled clients
				//
				//
				// # Use openssl to check the SNI certificate
				// https://www.claudiokuenzler.com/blog/693/curious-case-of-curl-ssl-tls-sni-http-host-header
				// https://stackoverflow.com/questions/3220419/openssl-s-client-using-a-proxy
				// i.e.
				// 		$ openssl s_client -proxy localhost:3129 -connect example.com
				//
				//		CONNECTED(00000005)
				//		2022/01/17 14:51:17 Connot support non-SNI enabled clients
				//
				// # SNI for a proxy
				// https://www.agwa.name/blog/post/writing_an_sni_proxy_in_go
				// The SNI value works well with split-horizon DNS where the proxy see the backend's private IP address
				// and external clients see the proxy's public IP address.
				// > a proxy server can read the server name and use it to decide where to route the connection,
				// > without having to decrypt the connection
				//
				// # SNI in the crypto/tls package
				// https://pkg.go.dev/crypto/tls#ClientHelloInfo
				// > ClientHelloInfo.ServerName indicates the name of the server requested by the client
				// > in order to support virtual hosting (only set if the client is using SNI)
				// >
				// > ClientHelloInfo.GetCertificate will only be called
				// > if the client supplies SNI information or if certificates is empty.
				//
				// # The term "Virtual hosting" in the field of HTTP
				// https://en.wikipedia.org/wiki/Virtual_hosting
				// Virtual hosting is a method for hosting multiple domain name on a single server.
				log.Printf("Connot support non-SNI enabled clients")
				return
			}
			// TODO: what kind of design model?
			connectReq := &http.Request{
				Method: http.MethodConnect,
				URL: &url.URL{
					// TODO: Opaque...?
					Opaque: tlsConn.Host(),
					// TODO: I don't know JoinHostPort
					// TODO: 443 is enough?
					Host: net.JoinHostPort(tlsConn.Host(), "443"),
				},
				Host:       tlsConn.Host(),
				Header:     make(http.Header),
				RemoteAddr: c.RemoteAddr().String(),
			}

			resp := dumbResponseWriter{tlsConn}
			// TODO: proxy.ServeHTTP directly use!
			proxy.ServeHTTP(resp, connectReq)
		}(c)
	}
}

func dial(ctx context.Context, proxy *goproxy.ProxyHttpServer, network, addr string) (c net.Conn, err error) {
	// Tr: HTTP Transport
	// TODO: Understand HTTP transport deeply
	if proxy.Tr.DialContext != nil {
		// TODO: use DialContext instead
		// TODO: make a pull request to use a context
		return proxy.Tr.DialContext(ctx, network, addr)
	}
	// TODO: What's difference between proxy.Tr.Dial and net.Dial
	var d net.Dialer
	return d.DialContext(ctx, network, addr)
}

// TODO: what's it?
func connectDial(ctx context.Context, proxy *goproxy.ProxyHttpServer, network, addr string) (c net.Conn, err error) {
	// ConnectDial will be used to create TCP connections for CONNECT requests if nil Tr.Dial will be used
	if proxy.ConnectDial == nil {
		return dial(ctx, proxy, network, addr)
	}
	return proxy.ConnectDial(network, addr)
}

// TODO: what is that? dumb?
type dumbResponseWriter struct {
	net.Conn
}

// TODO: Header is ok to panic?
func (d dumbResponseWriter) Header() http.Header {
	panic("Header() should not be called on this ResponseWriter")
}

func (d dumbResponseWriter) Write(buf []byte) (int, error) {
	if bytes.Equal(buf, []byte("HTTP/1.0 200 OK\r\n\r\n")) {
		// throw away the HTTP OK response from the faux CONNECT request
		// TODO: what is the faux CONNECT request
		return len(buf), nil
	}
	return d.Conn.Write(buf)
}

func (d dumbResponseWriter) WriteHeader(code int) {
	panic("WriteHeader() should not be called on this ResponseWriter")
}

// TODO: What is Hijack, which interface it implements?
func (d dumbResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return d, bufio.NewReadWriter(bufio.NewReader(d), bufio.NewWriter(d)), nil
}

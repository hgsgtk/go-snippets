package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {
	addr := "127.0.0.1:8080"

	p := &Proxy{}
	log.Printf("Starting proxy server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, p))
}

type Proxy struct{}

// Proxy implements the http.Handler interface.
// https://pkg.go.dev/net/http#Handler.ServeHTTP
// Also, refer to these:
// - https://gist.github.com/yowu/f7dc34bd4736a65ff28d
// - https://github.com/elazarl/goproxy/blob/adb46da277acd7aea06aeb8b5a21ec6bef7fb247/https.go#L87
func (p *Proxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	// Even if it's a https connection, we can read the following data.
	log.Printf(
		"client: %s, method: %s, url: %s, host: %s",
		req.RemoteAddr, // "127.0.0.1:8080"
		req.Method,     // "CONNECT"
		req.URL,        // "//example.com:443"
		req.Host,       // "google.com:443"
	)
	log.Printf("headers: %#v", req.Header)

	// Hijack lets the caller take over the connection, and returns net.Conn.
	// https://pkg.go.dev/net/http?utm_source=gopls#Hijacker
	hjj, ok := w.(http.Hijacker)
	if !ok {
		log.Fatal("httpserver does not support hijacking")
	}

	// Hijack lets the caller take over the connection.
	// https://pkg.go.dev/net/http#Hijacker.Hijack
	// Hijack() (net.Conn, *bufio.ReadWriter, error)
	// - The returned bufio.Reader may contain unprocessed buffered data from the client
	proxyClient, _, err := hjj.Hijack()
	if err != nil {
		log.Fatalf("Cannot hijack connection: %v", err)
	}

	// Dial to the target connection.
	// Fixme: some request don't have its port, check it.
	// https://github.com/elazarl/goproxy/blob/adb46da277acd7aea06aeb8b5a21ec6bef7fb247/https.go#L106
	var dialer net.Dialer
	targetSiteConn, err := dialer.DialContext(ctx, "tcp", req.URL.Host)
	if err != nil {
		httpError(proxyClient, err)
		return
	}
	log.Printf(
		"target site remote address: %s",
		targetSiteConn.RemoteAddr(), // e.g. google.com -> "142.250.199.110:443"
	)

	// Why goproxy adds an additional header?
	// https://github.com/elazarl/goproxy/blob/adb46da277acd7aea06aeb8b5a21ec6bef7fb247/https.go#L115
	proxyClient.Write([]byte("HTTP/1.0 200 OK\r\n\r\n"))

	// TODO: write the request to the target site.
	targetTCP, targetOK := targetSiteConn.(halfClosable)
	proxyClientTCP, clientOK := proxyClient.(halfClosable)
	if targetOK && clientOK {
		// copy the request from src: proxyClientTCP to dst: targetTCP
		go copyAndClose(targetTCP, proxyClientTCP)
		// copy the response from src: targetTCP to dst: proxyClientTCP
		go copyAndClose(proxyClientTCP, targetTCP)
	} else {
		// TODO: when scenario comes here?
		go func() {
			var wg sync.WaitGroup
			wg.Add(2)
			go copyOrWarn(targetSiteConn, proxyClient, &wg)
			go copyOrWarn(proxyClient, targetSiteConn, &wg)
			wg.Wait()
			proxyClient.Close()
			targetSiteConn.Close()
		}()
	}
}

// TODO: why io.WriteCloser
func httpError(w io.WriteCloser, err error) {
	if _, err := io.WriteString(w, "HTTP/1.1 502 Bad Gateway\r\n\r\n"); err != nil {
		log.Printf("Error responding to client: %v", err)
	}
	if err := w.Close(); err != nil {
		log.Printf("Error closing client connection: %v", err)
	}
}

// Why it is needed?
// https://github.com/elazarl/goproxy/blob/adb46da277acd7aea06aeb8b5a21ec6bef7fb247/https.go#L71
type halfClosable interface {
	net.Conn
	CloseWrite() error
	CloseRead() error
}

func copyAndClose(dst, src halfClosable) {
	// Copy copies from src to dst until either EOF is reached on src or an error occurs..
	// https://pkg.go.dev/io#Copy
	// If src implements the WriterTo interface, the copy is implemented by calling src.WriteTo(dst). Otherwise, if dst implements the ReaderFrom interface,
	// the copy is implemented by calling dst.ReadFrom(src).
	//
	// io.ReaderFrom interface.
	// https://pkg.go.dev/io#ReaderFrom
	// > ReadFrom reads data from r until EOF or error. The return value n is the number of bytes read. Any error except EOF encountered during the read is also returned.
	//
	// func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)
	// https://pkg.go.dev/net#TCPConn.ReadFrom
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("Error copying to client: %v", err)
	}

	dst.CloseWrite()
	src.CloseRead()
}

func copyOrWarn(dst io.Writer, src io.Reader, wg *sync.WaitGroup) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("Error copying to client: %v", err)
	}
	wg.Done()
}

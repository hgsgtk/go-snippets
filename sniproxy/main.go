// Ref: https://www.agwa.name/blog/post/writing_an_sni_proxy_in_go
package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":44433")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listening on %s\n", listener.Addr().String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("accepting connection from %s\n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	if err := clientConn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		log.Print(err)
		return
	}
	clientHello, clientReader, err := peekClientHello(clientConn)
	if err != nil {
		log.Print(err)
		return
	}
	// A zero value for t means Read will not time out.
	if err := clientConn.SetReadDeadline(time.Time{}); err != nil {
		log.Print(err)
		return
	}

	if !strings.HasSuffix(clientHello.ServerName, "example.com") {
		log.Print("Blocking connection to unauthorized backend")
		return
	}

	backendConn, err := net.DialTimeout(
		"tcp",
		net.JoinHostPort(clientHello.ServerName, "443"),
		5*time.Second,
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer backendConn.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		io.Copy(clientConn, backendConn)
		clientConn.(*net.TCPConn).CloseWrite()
		wg.Done()
	}()
	go func() {
		io.Copy(backendConn, clientReader)
		backendConn.(*net.TCPConn).CloseWrite()
		wg.Done()
	}()

	wg.Wait()
}

func peekClientHello(reader io.Reader) (*tls.ClientHelloInfo, io.Reader, error) {
	peekedBytes := new(bytes.Buffer)
	hello, err := readClientHello(io.TeeReader(reader, peekedBytes))
	if err != nil {
		return nil, nil, err
	}
	return hello, io.MultiReader(peekedBytes, reader), nil
}

func readClientHello(reader io.Reader) (*tls.ClientHelloInfo, error) {
	var hello *tls.ClientHelloInfo

	err := tls.Server(&readOnlyConn{reader: reader}, &tls.Config{
		// GetConfigForClient is called after a ClientHello is received from a client.
		// https://pkg.go.dev/crypto/tls#Config.GetConfigForClient
		GetConfigForClient: func(argHello *tls.ClientHelloInfo) (*tls.Config, error) {
			hello = new(tls.ClientHelloInfo)
			*hello = *argHello
			return nil, nil
		},
	}).Handshake()

	if hello == nil {
		return nil, err
	}

	return hello, nil
}

// readOnlyConn implements net.Conn interface.
type readOnlyConn struct {
	reader io.Reader
}

// Read reads data from the connection.
// Read can be made to time out and return an error after a fixed
// time limit; see SetDeadline and SetReadDeadline.
func (r *readOnlyConn) Read(b []byte) (n int, err error) {
	// $ curl -Lv --proxy http://127.0.0.1:44433 https://example.com
	//
	// 	* Proxy CONNECT aborted
	// 	* CONNECT phase completed!
	// 	* Closing connection 0
	//	curl: (56) Proxy CONNECT aborted
	//
	// It may causes the error - first record does not look like a TLS handshake
	// https://cs.opensource.google/go/go/+/refs/tags/go1.17.6:src/crypto/tls/conn.go;l=643
	// > First message, be extra suspicious: this might not be a TLS client. Bail out before reading a full 'body', if possible.
	// > The current max version is 3.3 so if the version is >= 16.0, it's probably not real.
	fmt.Printf("reading: %s\n", string(b)) // "": empty bytes.
	return r.reader.Read(b)
}

// Write writes data to the connection.
// Write can be made to time out and return an error after a fixed
// time limit; see SetDeadline and SetWriteDeadline.
func (r *readOnlyConn) Write(b []byte) (n int, err error) {
	panic("not implemented") // TODO: Implement
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (r *readOnlyConn) Close() error {
	panic("not implemented") // TODO: Implement
}

// LocalAddr returns the local network address.
func (r *readOnlyConn) LocalAddr() net.Addr {
	panic("not implemented") // TODO: Implement
}

// RemoteAddr returns the remote network address.
func (r *readOnlyConn) RemoteAddr() net.Addr {
	panic("not implemented") // TODO: Implement
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail instead of blocking. The deadline applies to all future
// and pending I/O, not just the immediately following call to
// Read or Write. After a deadline has been exceeded, the
// connection can be refreshed by setting a deadline in the future.
//
// If the deadline is exceeded a call to Read or Write or to other
// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
// The error's Timeout method will return true, but note that there
// are other possible errors for which the Timeout method will
// return true even if the deadline has not been exceeded.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
func (r *readOnlyConn) SetDeadline(t time.Time) error {
	panic("not implemented") // TODO: Implement
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (r *readOnlyConn) SetReadDeadline(t time.Time) error {
	panic("not implemented") // TODO: Implement
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (r *readOnlyConn) SetWriteDeadline(t time.Time) error {
	panic("not implemented") // TODO: Implement
}

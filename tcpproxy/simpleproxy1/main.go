// This is copying https://gist.github.com/vmihailenco/1380352 to lean tcp proxy in Go
package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
)

var lAddr string = "localhost:0" // a port number is automatically chosen
var rAddr string = "localhost:80"

func main() {

	fmt.Printf("Listening: %s\nProxing %s\n", lAddr, rAddr)

	// ResolveTCPAddr returns an address of TCP end point.
	// See https://pkg.go.dev/net#ResolveTCPAddr in more detail.
	// The first argument `network` could be tcp, tcp4, tcp6/
	addr, err := net.ResolveTCPAddr("tcp", lAddr)
	if err != nil {
		panic(err)
	}

	// ListenTCP acts like Listen for TCP networks.
	// See https://pkg.go.dev/net#ListenTCP in more detail.
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	pending := make(chan *net.TCPConn)
	complete := make(chan *net.TCPConn)

	for i := 0; i < 5; i++ {
		go handleConn(pending, complete)
	}
	go closeConn(complete)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
		pending <- conn
	}
}

func proxyConn(conn *net.TCPConn) error {
	rAddr, err := net.ResolveTCPAddr("tcp", rAddr)
	if err != nil {
		panic(err)
	}

	rConn, err := net.DialTCP("tcp", nil, rAddr)
	if err != nil {
		return err
	}
	defer rConn.Close()

	buf := &bytes.Buffer{}
	for {
		data := make([]byte, 256) // TODO: why 256?
		n, err := conn.Read(data)
		if err != nil {
			return err
		}
		buf.Write(data[:n])
		// TODO: what are these magic numbers?
		// 13: \r
		// 10: \n
		if data[0] == '\r' && data[1] == '\n' {
			break
		}
	}

	if _, err := rConn.Write(buf.Bytes()); err != nil {
		return err
	}
	// TODO: what's hex.Dump()
	log.Printf("send:\n%v", hex.Dump(buf.Bytes()))

	data := make([]byte, 1024)
	n, err := rConn.Read(data)
	if err != nil {
		if err != io.EOF {
			return err
		}
		log.Printf("received err: %v", err)
	}
	log.Printf("received:\n%v", hex.Dump(data[:n]))

	return nil
}

func handleConn(in <-chan *net.TCPConn, out chan<- *net.TCPConn) error {
	for conn := range in {
		if err := proxyConn(conn); err != nil {
			panic(err)
		}
		out <- conn
	}

	return nil
}

func closeConn(in <-chan *net.TCPConn) {
	for conn := range in {
		conn.Close()
	}
}

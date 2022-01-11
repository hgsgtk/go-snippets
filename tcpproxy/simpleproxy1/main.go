// This is copying https://gist.github.com/vmihailenco/1380352 to lean tcp proxy in Go
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var lAddr string = "localhost:9999" // if 0, a port number is automatically chosen
var rAddr string = "localhost:80"

// TODO: clear my mind by creating the server architecture diagram
// TODO: need tcp server for testing
func main() {
	fmt.Printf("Listening: %s\nProxing %s\n", lAddr, rAddr)

	// ListenTCP acts like Listen for TCP networks.
	// See https://pkg.go.dev/net#ListenTCP in more detail.
	listener, err := net.Listen("tcp", lAddr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection", err)
			continue
		}
		go func() {
			conn2, err := net.Dial("tcp", lAddr)
			if err != nil {
				log.Println("error dialing remote addr", err)
				return
			}
			// TODO: What is that?
			go io.Copy(conn2, conn)
			io.Copy(conn, conn2)
			conn2.Close()
			conn.Close()
		}()
	}
}

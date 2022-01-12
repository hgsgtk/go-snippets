// It is copying: https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go
package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections.
	addr := "localhost:8888"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Printf("Listening on %s\n", addr)

	for {
		// Listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		// Handle connections in a new goroutine
		go func(conn net.Conn) {
			buf := make([]byte, 1024)
			len, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("Error reading: %#v\n", err)
				return
			}
			fmt.Printf("Message received: %s\n", string(buf[:len]))

			conn.Write([]byte("Message received.\n"))
			conn.Close()
		}(conn)
	}
}

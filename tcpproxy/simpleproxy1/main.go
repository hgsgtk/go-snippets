// Ref: https://gist.github.com/vmihailenco/1380352#gistcomment-3731994
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	remoteAddr := make(chan string)
	// The remote address launched in another process.
	go runRemoteServer(remoteAddr)

	addr := <-remoteAddr
	runProxyServer(addr)
}

func runProxyServer(remoteAddr string) string {
	// Listen for incoming connection
	// if 0, a port number is automatically chosen
	localAddr := "localhost:0"
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		panic(err)
	}
	host, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"Proxy server listening on host: %s, port: %s. Proxing to %s\n",
		host,
		port,
		remoteAddr,
	)

	for {
		localConn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection", err)
			continue
		}
		// RemoteAddr returns the remote network address.
		fmt.Printf("Proxy server accepts new connection: %s\n", localConn.RemoteAddr())

		go func() {
			defer localConn.Close()

			// Proxy remote address
			remoteConn, err := net.Dial("tcp", remoteAddr)
			if err != nil {
				log.Println("error dialing remote addr", err)
				return
			}
			defer remoteConn.Close()

			// copy local connection to remote connection
			go io.Copy(remoteConn, localConn)
			// copy remote connection to local connection
			io.Copy(localConn, remoteConn)
		}()
	}
}

func runRemoteServer(serverStarted chan<- string) string {
	// Listen for incoming connections.
	addr := "127.0.0.1:" // a port number will be automatically chosen
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	host, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Remote server is listening on host: %s, port: %s\n", host, port)
	serverStarted <- l.Addr().String()

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
			fmt.Printf("Remote server received a message: %s\n", string(buf[:len]))

			conn.Write([]byte("Message received.\n"))
			conn.Close()
		}(conn)
	}
}

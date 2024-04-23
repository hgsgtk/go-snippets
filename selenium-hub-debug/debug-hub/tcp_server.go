package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func startTCPServer(port string) {
	// Specify the server address and port
	addr := fmt.Sprintf("localhost:%s", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}
	defer listener.Close()
	log.Printf("Server listening on %s", addr)

	for {
		// Wait for a connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// Handle the connection in a new goroutine.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Connected to %s", conn.RemoteAddr().String())

	// Create a new scanner to read from the connection
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("Received from %s: %s", conn.RemoteAddr().String(), text)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from %s: %v", conn.RemoteAddr().String(), err)
	} else {
		log.Printf("Connection closed by client: %s", conn.RemoteAddr().String())
	}
}

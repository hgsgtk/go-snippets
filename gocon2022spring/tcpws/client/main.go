package main

import (
	"io"
	"log"
	"net"

	"github.com/gorilla/websocket"
	"github.com/hgsgtk/go-snippets/gocon2022spring/tcpws"
)

func main() {
	wsConn, err := websocketHandshake()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connected to the WebSocket server")

	dest, err := wsConn.ReadHandshakeRequest()
	if err != nil {
		log.Fatalf("failed to read handshake request: %v", err)
	}

	tcpConn, err := tcpHandshake(dest)
	if err != nil {
		log.Fatalf("failed to connect to the destination: %v", err)
	}
	wrapTcpConn := tcpws.NewHijackedConn(tcpConn, nil)

	if err := wsConn.WriteHandshakeCompleted(); err != nil {
		log.Fatalf("failed to send handshake completed: %v", err)
	}

	log.Printf("connected to the destination: %s", dest)

	go io.Copy(wrapTcpConn, wsConn)
	go io.Copy(wsConn, wrapTcpConn)

	select {}
}

func websocketHandshake() (*tcpws.WrapWSConn, error) {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/websocket", nil)
	if err != nil {
		return nil, err
	}

	return &tcpws.WrapWSConn{RawConn: conn}, nil
}

func tcpHandshake(dest string) (net.Conn, error) {
	conn, err := net.Dial("tcp", dest)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

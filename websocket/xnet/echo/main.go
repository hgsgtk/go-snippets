// https://pkg.go.dev/golang.org/x/net/websocket
package main

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// https://pkg.go.dev/golang.org/x/net/websocket#example-Handler
// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	log.Fatal(http.ListenAndServe(":12345", nil))
}

package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hgsgtk/go-snippets/gocon2022spring/tcpws"
)

func main() {
	s := &Server{}
	http.ListenAndServe(":8080", s)
}

type Server struct {
	wsConn *tcpws.WrapWSConn
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/websocket" {
		s.Handshake(w, r)
		return
	}

	s.Proxy(w, r)
}

func (s *Server) Handshake(w http.ResponseWriter, r *http.Request) {
	wsConn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	s.wsConn = &tcpws.WrapWSConn{RawConn: wsConn}
	log.Printf("connected to the WebSocket client")
}

func (s *Server) Proxy(w http.ResponseWriter, r *http.Request) {
	if s.wsConn == nil {
		http.Error(w, "Not connected", http.StatusBadRequest)
		return
	}

	log.Printf("proxy to the destination: %s", r.Host)

	if err := s.wsConn.WriteHandshakeRequest(r.Host); err != nil {
		http.Error(w, "Could not send handshake request", http.StatusBadRequest)
		return
	}

	if err := s.wsConn.IsHandshaked(); err != nil {
		http.Error(w, "Failed to handshake", http.StatusBadRequest)
		return
	}

	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	conn, bufRw, err := hj.Hijack()
	if err != nil {
		http.Error(w, "Could not hijack connection", http.StatusInternalServerError)
		return
	}
	hjConn := tcpws.NewHijackedConn(conn, bufRw.Reader)

	res := []byte("HTTP/1.0 200 Connection established\r\n\r\n")
	if _, err := hjConn.Write(res); err != nil {
		http.Error(w, "Could not send handshake response", http.StatusInternalServerError)
		return
	}

	go io.Copy(hjConn, s.wsConn)
	go io.Copy(s.wsConn, hjConn)
}

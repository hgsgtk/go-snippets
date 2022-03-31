package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		hijacked, ok := w.(http.Hijacker)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tcpConn, _, err := hijacked.Hijack()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer tcpConn.Close()

		if _, err := tcpConn.Write([]byte("HTTP/1.0 200 Connection established\r\n\r\n")); err != nil {
			tcpConn.Close()
			return
		}

		<-ctx.Done() // Not coming.
		fmt.Print("context Done.")
	}

	r := http.NewServeMux()
	r.HandleFunc("/hijack", h)
	server := &http.Server{
		Handler: r,
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(l); err != nil {
		panic(err)
	}
}

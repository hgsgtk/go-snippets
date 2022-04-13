package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

type hijackHandler struct{}

func (h *hijackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	resp := "HTTP/1.0 200 OK\r\n\r\nGopher at your service\r\n"
	bufrw.WriteString(resp)
	bufrw.Flush()
}

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Fprintf(os.Stdout, "server on %q", l.Addr().String())
	if err := http.Serve(l, &hijackHandler{}); err != nil {
		panic(err)
	}
}

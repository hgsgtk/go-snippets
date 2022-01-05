package main

import (
	"fmt"
	"log"
	"net/http"
)

type ProxyServer struct{}

func (p *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func main() {
	mux := http.NewServeMux()
	// $ curl -x http://127.0.0.1:8080/connect http://example.com
	// proxy customized
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "proxy customized\n")
	})
	// $ curl http://127.0.0.1:8080/connect
	// connected
	mux.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "connected\n")
	})
	// HTTPS is not supported, how...?
	// $ curl -x http://127.0.0.1:8080/ https://example.com
	// curl: (56) Received HTTP code 404 from proxy after CONNECT

	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

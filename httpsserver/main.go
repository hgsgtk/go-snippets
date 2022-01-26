package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	go runTLSServer()

	// Wait until the server starts
	time.Sleep(1 * time.Second)
	execClient()
}

func runTLSServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	log.Fatal(http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", nil))
}

func execClient() {
	// https://doc.xuwenliang.com/docs/go/1397
	crt, err := os.ReadFile("localhost.crt")
	if err != nil {
		log.Fatalf("read certificate file: %v", err)
	}

	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(crt); !ok {
		log.Fatalf("append certs: %v", err)
	}

	cert, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
	if err != nil {
		log.Fatalf("load x509 key pair: %v", err)
	}

	// https://github.com/jcbsmpsn/golang-https-example/blob/6fa58aeeea418166caf61094e7207a92353d9623/https_client.go
	tlsConf := &tls.Config{
		RootCAs:      roots,
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{TLSClientConfig: tlsConf}
	client := &http.Client{Transport: tr}

	res, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatalf("error new request: %v", err)
	}

	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading: %v", err)
	}

	log.Printf(string(resBytes))
}

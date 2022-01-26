package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServeTLS(":9000", "localhost.crt", "localhost.key", nil))

}

func execClient() {
	// https://doc.xuwenliang.com/docs/go/1397
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:9000", conf)
	if err != nil {
		log.Fatalf("dial error: %v", err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("Hello"))
	if err != nil {
		log.Fatalf("writing error: %v", err)
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatalf("reading error: %v", err)
	}

	log.Printf(string(buf[:n]))
}

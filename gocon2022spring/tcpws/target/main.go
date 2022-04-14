package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
)

func main() {
	h := func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("failed to read body: %v", err)
		}
		log.Printf("body: %s", body)

		w.Write([]byte("hello"))
	}

	ts := httptest.NewUnstartedServer(http.HandlerFunc(h))
	ts.StartTLS()
	log.Printf("server on %q", ts.Listener.Addr().String())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Print("stopping")
}

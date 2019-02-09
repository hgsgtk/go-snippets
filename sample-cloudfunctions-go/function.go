package sample_cloudfunctions_go

import "net/http"

func Hello(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World from Mac"
	w.Write([]byte(msg))
}

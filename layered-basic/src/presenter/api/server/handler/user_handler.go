package handler

import (
	"fmt"
	"net/http"
)

// IndexHandler return hello message
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, http server.")
}

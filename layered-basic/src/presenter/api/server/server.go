package server

import (
	"net/http"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"

	"github.com/Khigashiguchi/go-snippets/layered-basic/src/presenter/api/server/handler"
)

// Run run http server
func Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/users/", handler.IndexHandler)

	http.Handle("/", httpauth.SimpleBasicAuth("test", "test")(r))
	return http.ListenAndServe(":8080", nil)
}

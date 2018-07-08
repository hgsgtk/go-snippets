package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	// See template.go
	indexTmpl      = parseTemplate("index.html")
	adminIndexTmpl = parseTemplate("admin_index.html")
)

func main() {
	registerHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerHandlers() {
	r := mux.NewRouter()

	r.Methods("GET").Path("/").Handler(appHandler(indexHandler))

	// The following handlers are defined in auth.go
	r.Methods("GET").Path("/login").Handler(appHandler(loginHandler))
	r.Methods("GET").Path("/oauth2callback").Handler(appHandler(oauthCallbackHandler))
	r.Methods("POST").Path("/logout").Handler(appHandler(oauthCallbackHandler))

	// The following handlers are defined in admin.go
	r.Methods("GET").Path("/admin/").Handler(appHandler(adminIndexHandler))

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}

// http://blog.golang.org/error-handling-and-go
type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v", e.Code, e.Message, e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}
func indexHandler(w http.ResponseWriter, r *http.Request) *appError {
	return indexTmpl.Execute(w, r, nil)
}

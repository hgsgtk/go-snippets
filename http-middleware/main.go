package main

import (
	"net/http"
	"fmt"
	"os"
	"time"
	"github.com/gorilla/mux"
	"log"
)

func helloHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world.")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		t := t2.Sub(t1)
		fmt.Fprintf(os.Stdout, "[%s] %s %s", r.Method, r.URL, t.String())
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(os.Stdout, "failed to unexpected reason: %+v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}


func appLoggingMiddleware(logger *log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			next.ServeHTTP(w, r)
			t2 := time.Now()
			t := t2.Sub(t1)
			logger.Printf("[%s] %s %s", r.Method, r.URL, t.String())
		})
	}
}

type authenticationMiddleware struct {
	tokenUsers map[string]string
}

func (amw *authenticationMiddleware) Populate() {
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["aaaaaaaa"] = "userA"
	amw.tokenUsers["05f717e5"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			fmt.Fprintf(os.Stdout, "Authenticate user %s\n", user)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func main() {
	// only standard package
	//mux := http.NewServeMux()
	//mux.HandleFunc("/hello", recoveryMiddleware(loggingMiddleware(helloHandler)))
	// use gorilla/mux
	logger := log.New(os.Stdout, "", 0)

	router := mux.NewRouter()

	amw := authenticationMiddleware{}
	amw.Populate()

	r := router.PathPrefix("/api").Subrouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.Use(recoveryMiddleware, loggingMiddleware, amw.Middleware, appLoggingMiddleware(logger))
	fmt.Fprint(os.Stdout, "start to http server.\n")
	if err := http.ListenAndServe(":3000", router); err != nil {
		fmt.Fprintf(os.Stderr, "faild to start http server: %s", err)
	}
}


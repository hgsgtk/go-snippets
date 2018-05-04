package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type Book struct {
    ID string `json:"id"`
    Isbn string `json:"isbn"`
    Title string `json:"title"`
    Author *Author `json:"author"`
}

type Author struct {
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range books {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = strconv.Itoa(rand.Intn(1000000000))
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"] {
            books = append(books[:index], books[index+1:]...)
            var book Book
            _ = json.NewDecoder(r.Body).Decode(&book)
            book.ID = params["id"]
            books = append(books, book)
            json.NewEncoder(w).Encode(book)
            return
        }
    }
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"] {
            books = append(books[:index], books[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(books)
    }
}

func main() {
    r := mux.NewRouter()

    books = append(books, Book{ID: "1", Isbn: "1234567", Title: "Book One", Author: &Author{Firstname: "Taro", Lastname: "Suzuki"}})
    books = append(books, Book{ID: "2", Isbn: "8901234", Title: "Book Two", Author: &Author{Firstname: "Hanako", Lastname: "Yamada"}})

    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/books", createBook).Methods("POST")
    r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", r))
}

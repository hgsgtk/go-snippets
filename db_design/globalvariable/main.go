package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Khigashiguchi/go-snippets/db_design/globalvariable/models"
)

func init() {
	var err error
	err = models.InitDB("sqlite3", "./simple.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/create", booksCreate)
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	bks, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, %d円\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, "paramter: isbn is necessary.", http.StatusBadRequest)
		return
	}

	bk, err := models.GetBookByIsbn(isbn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "%s, %s, %s, %d円\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
}

func booksCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	isbn := r.FormValue("isbn")
	title := r.FormValue("title")
	author := r.FormValue("author")
	if isbn == "" || title == "" || author == "" {
		http.Error(w, "paramter: isbn, title, author is necessary.", http.StatusBadRequest)
		return
	}
	price, err := strconv.ParseInt(r.FormValue("price"), 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rowAffected, err := models.CreateBook(isbn, title, author, price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", isbn, rowAffected)
}

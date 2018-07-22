package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Khigashiguchi/go-snippets/db_design/usingctx/models"
)

// ContextInjector represents context injection struct.
type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ci.h.ServeHTTP(w, r.WithContext(ci.ctx))
}

func main() {
	db, err := models.NewDB("sqlite3", "./sample.sqlite3")
	if err != nil {
		log.Panic(err)
	}
	ctx := context.WithValue(context.Background(), "db", db)

	http.Handle("/books", &ContextInjector{ctx, http.HandlerFunc(booksIndex)})
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	db, ok := r.Context().Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "could not get database pool from context", http.StatusInternalServerError)
		return
	}

	bks, err := models.GetAllBooks(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, %d円\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

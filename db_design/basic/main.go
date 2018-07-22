package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Book represents a structure of book data.
type Book struct {
	isbn   string
	title  string
	author string
	price  int
}

func main() {
	db, err := sql.Open("sqlite3", "./sample.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
		bks = append(bks, bk)
	}
	if err := rows.Err(); err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	for _, bk := range bks {
		fmt.Fprintf(os.Stdout, "%s, %s, %s, %d円\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

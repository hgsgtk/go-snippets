package models

import (
	"database/sql"
	"errors"
)

// Book represents a structure of book data.
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  int
}

// GetAllBooks execute to fecth all records in bools.
func GetAllBooks() ([]*Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

// GetBookByIsbn execute to fetch one record by isbn.
func GetBookByIsbn(isbn string) (*Book, error) {
	row := db.QueryRow("SELECT * FROM books WHERE isbn = ?", isbn)

	bk := new(Book)
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err == sql.ErrNoRows {
		return nil, errors.New("not found")
	} else if err != nil {
		return nil, err
	}
	return bk, nil
}

// CreateBook execute to save one record in books.
func CreateBook(isbn string, title string, author string, price int64) (int64, error) {
	result, err := db.Exec("INSERT INTO books VALUES(?, ?, ?, ?)", isbn, title, author, price)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowAffected, nil
}

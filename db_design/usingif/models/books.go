package models

// Book represents a structure of book data.
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  int
}

// GetAllBooks execute to fecth all records in bools.
func (db *DB) GetAllBooks() ([]*Book, error) {
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

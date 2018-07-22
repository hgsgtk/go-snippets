package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Khigashiguchi/go-snippets/db_design/usingif/models"
)

type mockDB struct{}

func (mdb *mockDB) GetAllBooks() ([]*models.Book, error) {
	bks := make([]*models.Book, 0)
	bks = append(bks, &models.Book{"978-1503261969", "Emma", "Jayne Austen", 10000})
	bks = append(bks, &models.Book{"978-1505255607", "The Time Machine", "H. G. Wells", 5000})
	return bks, nil
}

func TestBooksIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.booksIndex).ServeHTTP(rec, req)

	expected := "978-1503261969, Emma, Jayne Austen, 10000円\n978-1505255607, The Time Machine, H. G. Wells, 5000円\n"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

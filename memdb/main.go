package main

import (
	"fmt"

	"github.com/hashicorp/go-memdb"
)

func main() {
	type PersonID int

	// Create a sample struct
	type Person struct {
		ID    PersonID
		Email string
		Name  string
		Age   int
	}

	type Devise struct {
		PersonID PersonID
	}

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"person": {
				Name: "person",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
					"age": {
						Name:    "age",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Age"},
					},
				},
			},
			"devise": {
				Name: "devise",
				Indexes: map[string]*memdb.IndexSchema{
					"person_id": {
						Name:    "person_id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "PersonID"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// Create a write transaction
	txn := db.Txn(true)

	// Insert some people
	people := []*Person{
		&Person{1, "joe@aol.com", "Joe", 30},
		&Person{2, "lucy@aol.com", "Lucy", 35},
		&Person{3, "tariq@aol.com", "Tariq", 21},
		&Person{4, "dorothy@aol.com", "Dorothy", 53},
	}
	for _, p := range people {
		if err := txn.Insert("person", p); err != nil {
			panic(err)
		}
	}
	devise := []*Devise{
		&Devise{1},
		&Devise{3},
	}
	for _, d := range devise {
		if err := txn.Insert("devise", d); err != nil {
			panic(err)
		}
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = db.Txn(false)
	defer txn.Abort()

	// Lookup by email
	raw, err := txn.First("person", "id", "joe@aol.com")
	if err != nil {
		panic(err)
	}

	// Say hi!
	fmt.Printf("Hello %s!\n", raw.(*Person).Name)

	raw, err = txn.First("devise", "person_id", 1)
	if err != nil {

	}

	// List all the people
	it, err := txn.Get("person", "id")
	if err != nil {
		panic(err)
	}

	fmt.Println("All the people:")
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Person)
		fmt.Printf("  %s\n", p.Name)
	}

	// Range scan over people with ages between 25 and 35 inclusive
	it, err = txn.LowerBound("person", "age", 25)
	if err != nil {
		panic(err)
	}

	fmt.Println("People aged 25 - 35:")
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Person)
		if p.Age > 35 {
			break
		}
		fmt.Printf("  %s is aged %d\n", p.Name, p.Age)
	}
	// Output:
	// Hello Joe!
	// All the people:
	//   Dorothy
	//   Joe
	//   Lucy
	//   Tariq
	// People aged 25 - 35:
	//   Joe is aged 30
	//   Lucy is aged 35
}

package main

import (
	"log"
	"os"
)

func main() {
	dbHost := os.Getenv("MAIN_DB_HOST")
	if len(dbHost) == 0 {
		log.Fatal("missing env MAIN_DB_HOST")
	}
	//  ... something
}

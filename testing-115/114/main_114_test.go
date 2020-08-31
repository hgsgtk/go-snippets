package main_test

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	shutdown := setupDB()
	c := m.Run()
	shutdown()
	os.Exit(c)
}

func setupDB() func() {
	return func() {
		log.Print("called")
		// shutdown
	}
}

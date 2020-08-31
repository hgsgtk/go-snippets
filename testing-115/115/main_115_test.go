package main_test

import (
	"testing"
)

func TestMain(m *testing.M) {
	shutdown := setupDB()

	m.Run()

	shutdown()
}

func setupDB() func() {
	return func() {
		// shutdown
	}
}

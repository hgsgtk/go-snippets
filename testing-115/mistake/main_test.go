package main_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var exitCode int
	defer os.Exit(exitCode)

	setup()
	defer teardown()
	exitCode = m.Run()
}

func TestHoge(t *testing.T) {
	t.Fail()
}

func setup() {
	//
}

func teardown() {
	//
}

package main_test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	defer teardown()

	var exitCode int
	exitCall := func() {
		fmt.Println("call exit by defer()")
		os.Exit(exitCode)
	}
	defer exitCall()
	_ = exitCode
	exitCode = m.Run()
}

func TestHoge(t *testing.T) {
	t.Fail()
}

func setup() {
	//
}

func teardown() {
	panic("call teardown()")
	//
}

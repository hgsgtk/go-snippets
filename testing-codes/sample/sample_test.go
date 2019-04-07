package sample_test

import (
	"github.com/hgsgtk/go-snippets/testing-codes/sample"
	"testing"
)

func TestName(t *testing.T) {
	// ...
}

func TestSayHello(t *testing.T) {
	want := "hello"
	// SayHelloの戻り値が期待値とことなる場合エラーとして処理する
	if got := sample.SayHello(); got != want {
		t.Errorf("SayHello() = %#v, want %#v", got, want)
	}
}

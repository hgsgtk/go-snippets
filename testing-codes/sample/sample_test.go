package sample_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hgsgtk/go-snippets/testing-codes/sample"
)

func TestName(t *testing.T) {
	// ...
}

func TestSayHello(t *testing.T) {
	want := "hello"
	// SayHelloの戻り値が期待値とことなる場合エラーとして処理する
	if got := sample.SayHello(); got != want {
		t.Errorf("SayHello() = %s, want %s", got, want)
	}
}

func TestGetNum(t *testing.T) {
	str := "7"
	got, err := sample.GetNum(str)
	if err != nil {
		t.Fatalf("GetNum(%s) caused unexpected error '%#v'", str, err)
	}
	want := 7
	if got != want {
		t.Errorf("GetNum(%s) = %d, want %d", str, got, want)
	}
}

func TestGetTomorrowUsingCmp(t *testing.T) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("Failed to load JST time Location")
	}
	tm := time.Date(2019, time.April, 14, 0, 0, 0, 0, jst)

	want := tm.AddDate(0, 0, 1)
	got := sample.GetTomorrow(tm)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("GetTomorrow() differs: (-want +got)\n%s", diff)
	}
}

package sample_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hgsgtk/go-snippets/testing-codes/testutil"

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

func TestInStatusListWorse(t *testing.T) {
	var x string
	var want bool

	x = "deleted" // 実は deleted というステータスもあった
	want = true
	if got := sample.InStatusList(x); got != want {
		t.Errorf("unexpected value %t", got)
	}
}

func TestInStatusListBetter(t *testing.T) {
	var x string
	var want bool

	x = "deleted"
	want = true
	if got := sample.InStatusList(x); got != want {
		t.Errorf("InStatusList(%s) = %t, want %t", x, got, want)
	}
}

func TestGetTomorrowUsingCmp(t *testing.T) {
	tm := time.Date(2019, time.April, 14, 0, 0, 0, 0, testutil.GetJstLocation(t))

	want := tm.AddDate(0, 0, 3)
	got := sample.GetTomorrow(tm)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("GetTomorrow() differs: (-got +want)\n%s", diff)
	}
}

func TestOkHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/sample", nil)

	sample.OkHandler(w, r)
	res := w.Result()
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll() caused unexpected error '%#v'", err)
	}
	const expected = "{\"status\":\"OK\"}\n"
	if got := string(b); got != expected {
		t.Errorf("OkHandler response = '%#v', want '%#v'", got, expected)
	}
}

package simple_test

import (
	"github.com/hgsgtk/go-snippets/testing-codes/testutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hgsgtk/go-snippets/testing-codes/simple"
)

func TestName(t *testing.T) {
	// ...
}

func TestSayHello(t *testing.T) {
	want := "hello"
	// SayHelloの戻り値が期待値とことなる場合エラーとして処理する
	if got := simple.SayHello(); got != want {
		t.Errorf("SayHello() = %s, want %s", got, want)
	}
}

func TestGetNum(t *testing.T) {
	str := "7"
	got, err := simple.GetNum(str)
	if err != nil {
		t.Fatalf("GetNum(%s) caused unexpected error '%#v'", str, err)
	}
	want := 7
	if got != want {
		t.Errorf("GetNum(%s) = %d, want %d", str, got, want)
	}
}

func TestInStatusListImproperErrorReport(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{
			arg:  "unknown1",
			want: false,
		},
		{
			arg:  "drafted",
			want: true,
		},
		{
			arg:  "unknown2",
			want: false,
		},
	}

	for _, test := range tests {
		if got := simple.InStatusList(test.arg); got != test.want {
			t.Errorf("unexpected value %t", got)
		}
	}
}

func TestInStatusListProperErrorReport(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{
			arg:  "unknown1",
			want: false,
		},
		{
			arg:  "drafted",
			want: true,
		},
		{
			arg:  "unknown2",
			want: false,
		},
	}

	for _, tt := range tests {
		if got := simple.InStatusList(tt.arg); got != tt.want {
			t.Errorf("InStatusList(`%s`) = %t, want %t", tt.arg, got, tt.want)
		}
	}
}

func TestInStatusListBetter(t *testing.T) {
	var x string
	var want bool

	x = "deleted"
	want = true
	if got := simple.InStatusList(x); got != want {
		t.Errorf("InStatusList(%s) = %t, want %t", x, got, want)
	}
}

func TestGetTomorrowUsingCmp(t *testing.T) {
	tm := time.Date(2019, time.April, 14, 0, 0, 0, 0,
		testutil.GetJstLocation(t))

	want := tm.AddDate(0, 0, 3)
	got := simple.GetTomorrow(tm)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("GetTomorrow() differs: (-got +want)\n%s", diff)
	}
}

func TestOkHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/simple", nil)

	simple.OkHandler(w, r)
	res := w.Result()
	defer res.Body.Close()

	// Check http response
	wantCode := http.StatusOK
	wantBody := "./testdata/okhandler.json.golden"
	testutil.AssertRespoonse(t, "OkHandler", res, wantCode, wantBody)
}

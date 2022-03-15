package main_test

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"
)

func TestHTTPServer(t *testing.T) {
	th := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("body message"))
	}
	ts := httptest.NewServer(http.HandlerFunc(th))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	bt, err := httputil.DumpResponse(res, true)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%q", string(bt))
}

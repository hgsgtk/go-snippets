package testutil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// AssertResponse assert response header and body.
func AssertResponse(
	t *testing.T, testTarget string,
	gotRes *http.Response,
	wantCode int, bodyFile string) {
	t.Helper()

	AssertResponseHeader(t, testTarget, gotRes, wantCode)
	AssertResponseBodyWithFile(t, testTarget, gotRes, bodyFile)
}

// AssertResponseHeader assert response header.
func AssertResponseHeader(
	t *testing.T, testTarget string,
	gotRes *http.Response,
	wantCode int) {
	t.Helper()

	// Check status code
	if got := gotRes.StatusCode; got != wantCode {
		t.Errorf("%s respond http status '%d', want '%d'", testTarget, got, wantCode)
	}
	// Check Content-Type
	wantType := "application/json; charset=utf-8"
	if got := gotRes.Header.Get("Content-Type"); got != wantType {
		t.Errorf("%s respond Content-Type '%s', want '%s'", testTarget, got, wantType)
	}
}

// AssertResponseBodyWithFile assert response body with test file.
func AssertResponseBodyWithFile(
	t *testing.T, testTarget string,
	gotRes *http.Response,
	bodyFile string) {
	t.Helper()

	rs := getStringFromTestFile(t, bodyFile)
	body, err := ioutil.ReadAll(gotRes.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll() got unexpected error %#v", err)
	}
	if diff := cmp.Diff(string(body), rs, getJSONCmpOption(t)); diff != "" {
		t.Errorf("%s's response body got differs: (-got +want)\n%s", testTarget, diff)
	}
}

func getJSONCmpOption(t *testing.T) cmp.Option {
	xform := cmp.Transformer("JSONcmp", func(s string) (m map[string]interface{}) {
		if err := json.Unmarshal([]byte(s), &m); err != nil {
			t.Fatalf("json.Unmarshal(%s) got unexpected error %#v", s, err)
		}
		return m
	})
	opt := cmp.FilterPath(func(p cmp.Path) bool {
		for _, ps := range p {
			if tr, ok := ps.(cmp.Transform); ok && tr.Option() == xform {
				return false
			}
		}
		return true
	}, xform)
	return opt
}

func getStringFromTestFile(t *testing.T, path string) string {
	t.Helper()

	bt, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("ioutil.ReadFile() got unexpected error while opening file %#v", err)
	}
	return string(bt)
}

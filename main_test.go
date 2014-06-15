package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(topHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/test")
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}

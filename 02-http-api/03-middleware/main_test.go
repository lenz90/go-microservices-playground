package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRec(t *testing.T) {
	h := chain(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { panic("x") }), rec(slog.New(slog.NewTextHandler(os.Stdout, nil))))
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/", nil))
	if rr.Code != 500 {
		t.Fatal()
	}
}

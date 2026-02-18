package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	rr := httptest.NewRecorder()
	newRouter().ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/v1/health", nil))
	if rr.Code != 200 {
		t.Fatal(rr.Code)
	}
}

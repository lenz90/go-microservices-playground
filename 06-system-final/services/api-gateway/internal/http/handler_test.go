package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	h, _ := New("http://localhost:18081", "http://localhost:18082")
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	h.Router().ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Fatalf("got %d", rr.Code)
	}
}

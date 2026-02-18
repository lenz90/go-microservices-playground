package httpx

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakePub struct{}

func (fakePub) UserCreated(context.Context, string, string) error { return nil }

func TestCreateUser(t *testing.T) {
	h := New(fakePub{}).Router()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"id":"u1","name":"Ada","email":"a@x.com"}`))
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		t.Fatalf("got %d", rr.Code)
	}
}

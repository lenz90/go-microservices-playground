package httpx

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"playground/06-system-final/services/payments-service/internal/storage"
)

type fakeStore struct{}

func (fakeStore) Create(context.Context, storage.Payment) error { return nil }

type fakePub struct{}

func (fakePub) PaymentCreated(context.Context, string, string, float64) error { return nil }

func TestCreatePayment(t *testing.T) {
	h := New(fakeStore{}, fakePub{}).Router()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/payments", strings.NewReader(`{"id":"p1","user_id":"u1","amount":10}`))
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		t.Fatalf("got %d", rr.Code)
	}
}

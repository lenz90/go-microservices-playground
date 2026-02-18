package httpx

import (
	"context"
	"encoding/json"
	"net/http"

	"playground/06-system-final/services/payments-service/internal/storage"
)

type EventPublisher interface {
	PaymentCreated(ctx context.Context, id, userID string, amount float64) error
}
type Store interface {
	Create(ctx context.Context, p storage.Payment) error
}

type Handler struct {
	store Store
	pub   EventPublisher
}

func New(store Store, pub EventPublisher) Handler { return Handler{store: store, pub: pub} }

func (h Handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte("ok")) })
	mux.HandleFunc("/payments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method", 405)
			return
		}
		var p storage.Payment
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if p.ID == "" || p.UserID == "" {
			http.Error(w, "id/user_id required", 400)
			return
		}
		if err := h.store.Create(r.Context(), p); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if err := h.pub.PaymentCreated(r.Context(), p.ID, p.UserID, p.Amount); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(p)
	})
	return mux
}

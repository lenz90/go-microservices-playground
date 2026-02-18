package httpx

import (
	"context"
	"encoding/json"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EventPublisher interface {
	UserCreated(ctx context.Context, id, email string) error
}

type Handler struct{ pub EventPublisher }

func New(pub EventPublisher) Handler { return Handler{pub: pub} }

func (h Handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte("ok")) })
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method", http.StatusMethodNotAllowed)
			return
		}
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if u.ID == "" || u.Email == "" {
			http.Error(w, "id/email required", http.StatusBadRequest)
			return
		}
		if err := h.pub.UserCreated(r.Context(), u.ID, u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(u)
	})
	return mux
}

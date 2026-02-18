package httptransport

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"playground/03-microservice-foundation/01-clean-arch/internal/domain"
	"playground/03-microservice-foundation/01-clean-arch/internal/usecase"
)

func Router(s usecase.Service) http.Handler {
	r := chi.NewRouter()
	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		var u domain.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if err := s.Create(r.Context(), u); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.WriteHeader(201)
	})
	return r
}

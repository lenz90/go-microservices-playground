package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
	"time"
)

type Product struct{ ID, Name string }
type Service struct {
	cache *redis.Client
	db    map[string]Product
	ttl   time.Duration
}

func (s Service) GetProduct(ctx context.Context, id string) (Product, error) {
	if raw, err := s.cache.Get(ctx, "product:"+id).Result(); err == nil {
		var p Product
		if json.Unmarshal([]byte(raw), &p) == nil {
			return p, nil
		}
	}
	time.Sleep(10 * time.Millisecond)
	p := s.db[id]
	b, _ := json.Marshal(p)
	_ = s.cache.Set(ctx, "product:"+id, b, s.ttl).Err()
	return p, nil
}
func main() {
	a := os.Getenv("REDIS_ADDR")
	if a == "" {
		a = "localhost:6379"
	}
	svc := Service{cache: redis.NewClient(&redis.Options{Addr: a}), db: map[string]Product{"p1": {"p1", "Keyboard"}}, ttl: time.Minute}
	r := chi.NewRouter()
	r.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		p, _ := svc.GetProduct(r.Context(), "p1")
		_ = json.NewEncoder(w).Encode(p)
	})
	_ = http.ListenAndServe(":8080", r)
}

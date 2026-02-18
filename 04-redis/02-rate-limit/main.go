package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"net/http"
	"os"
	"time"
)

type Limiter struct {
	rdb    *redis.Client
	limit  int
	window time.Duration
}

func (l Limiter) Allow(ctx context.Context, key string) (bool, error) {
	cnt, err := l.rdb.Incr(ctx, "rl:"+key).Result()
	if err != nil {
		return false, err
	}
	if cnt == 1 {
		_ = l.rdb.Expire(ctx, "rl:"+key, l.window).Err()
	}
	return cnt <= int64(l.limit), nil
}
func main() {
	a := os.Getenv("REDIS_ADDR")
	if a == "" {
		a = "localhost:6379"
	}
	lim := Limiter{rdb: redis.NewClient(&redis.Options{Addr: a}), limit: 5, window: time.Minute}
	r := chi.NewRouter()
	r.Use(func(n http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ok, _ := lim.Allow(req.Context(), "k")
			if !ok {
				http.Error(w, "rate", 429)
				return
			}
			n.ServeHTTP(w, req)
		})
	})
	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte("pong")) })
	_ = http.ListenAndServe(":8080", r)
}

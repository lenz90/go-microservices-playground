package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	s := Service{cache: redis.NewClient(&redis.Options{Addr: "x"}), db: map[string]Product{"p1": {"p1", "n"}}, ttl: time.Second}
	p, _ := s.GetProduct(context.Background(), "p1")
	if p.ID != "p1" {
		t.Fatal()
	}
}

package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestAllow(t *testing.T) {
	l := Limiter{rdb: redis.NewClient(&redis.Options{Addr: "x"}), limit: 2, window: time.Second}
	_, _ = l.Allow(context.Background(), "k")
	_, _ = l.Allow(context.Background(), "k")
	ok, _ := l.Allow(context.Background(), "k")
	if ok {
		t.Fatal()
	}
}

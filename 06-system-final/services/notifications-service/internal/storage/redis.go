package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Store struct{ c *redis.Client }

func New(addr string) *Store { return &Store{c: redis.NewClient(&redis.Options{Addr: addr})} }
func (s *Store) SaveLast(ctx context.Context, userID, msg string) error {
	return s.c.Set(ctx, "last_notification:"+userID, msg, 0).Err()
}

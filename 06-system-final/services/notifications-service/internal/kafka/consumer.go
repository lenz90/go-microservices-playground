package kafka

import (
	"context"
	"encoding/json"

	kgo "github.com/segmentio/kafka-go"
)

type Notifier interface {
	SaveLast(ctx context.Context, userID, msg string) error
}

type Consumer struct {
	usersR, paymentsR *kgo.Reader
	notifier          Notifier
}

func New(broker string, n Notifier) *Consumer {
	return &Consumer{
		usersR:    kgo.NewReader(kgo.ReaderConfig{Brokers: []string{broker}, Topic: "events.users", GroupID: "notifications-users"}),
		paymentsR: kgo.NewReader(kgo.ReaderConfig{Brokers: []string{broker}, Topic: "events.payments", GroupID: "notifications-payments"}),
		notifier:  n,
	}
}

func (c *Consumer) Close() error { _ = c.usersR.Close(); return c.paymentsR.Close() }

func extractUserID(b []byte) string {
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	if v, ok := m["user_id"].(string); ok {
		return v
	}
	return ""
}

func (c *Consumer) PollOnce(ctx context.Context) {
	if m, err := c.usersR.FetchMessage(ctx); err == nil {
		uid := extractUserID(m.Value)
		if uid != "" {
			_ = c.notifier.SaveLast(ctx, uid, string(m.Value))
		}
		_ = c.usersR.CommitMessages(ctx, m)
	}
	if m, err := c.paymentsR.FetchMessage(ctx); err == nil {
		uid := extractUserID(m.Value)
		if uid != "" {
			_ = c.notifier.SaveLast(ctx, uid, string(m.Value))
		}
		_ = c.paymentsR.CommitMessages(ctx, m)
	}
}

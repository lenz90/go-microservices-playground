package kafka

import (
	"context"
	"errors"
	"time"
)

type Header struct {
	Key   string
	Value []byte
}
type Message struct {
	Key, Value []byte
	Headers    []Header
	Offset     int64
}

type Writer struct {
	Addr         any
	Topic        string
	RequiredAcks int
}

func (w *Writer) WriteMessages(context.Context, ...Message) error { return nil }
func (w *Writer) Close() error                                    { return nil }

const RequireOne = 1

func TCP(b ...string) any { return b }

type Reader struct{}
type ReaderConfig struct {
	Brokers        []string
	Topic, GroupID string
}

func NewReader(ReaderConfig) *Reader { return &Reader{} }
func (r *Reader) ReadMessage(context.Context) (Message, error) {
	return Message{}, errors.New("no messages")
}
func (r *Reader) FetchMessage(ctx context.Context) (Message, error) {
	select {
	case <-ctx.Done():
		return Message{}, ctx.Err()
	case <-time.After(10 * time.Millisecond):
		return Message{}, errors.New("no messages")
	}
}
func (r *Reader) CommitMessages(context.Context, ...Message) error { return nil }
func (r *Reader) Close() error                                     { return nil }

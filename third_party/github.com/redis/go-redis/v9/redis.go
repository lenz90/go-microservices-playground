package redis

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Options struct{ Addr string }

type Client struct {
	mu   sync.Mutex
	data map[string]string
	n    map[string]int64
}

func NewClient(*Options) *Client { return &Client{data: map[string]string{}, n: map[string]int64{}} }

type StringCmd struct {
	v   string
	err error
}

func (c *StringCmd) Result() (string, error) { return c.v, c.err }

type StatusCmd struct{ err error }

func (c *StatusCmd) Err() error { return c.err }

type IntCmd struct {
	v   int64
	err error
}

func (c *IntCmd) Result() (int64, error) { return c.v, c.err }

func (c *Client) Get(_ context.Context, key string) *StringCmd {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.data[key]
	if !ok {
		return &StringCmd{"", fmt.Errorf("missing")}
	}
	return &StringCmd{v, nil}
}
func (c *Client) Set(_ context.Context, key string, value interface{}, _ time.Duration) *StatusCmd {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = fmt.Sprint(value)
	return &StatusCmd{}
}
func (c *Client) Incr(_ context.Context, key string) *IntCmd {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n[key]++
	return &IntCmd{v: c.n[key]}
}
func (c *Client) Expire(_ context.Context, key string, _ time.Duration) *BoolCmd {
	return &BoolCmd{v: true}
}

type BoolCmd struct {
	v   bool
	err error
}

func (c *BoolCmd) Err() error { return c.err }

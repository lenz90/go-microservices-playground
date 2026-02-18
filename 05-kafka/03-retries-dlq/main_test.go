package main

import (
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestRetryCount(t *testing.T) {
	if retryCount([]kafka.Header{{Key: "x-retry-count", Value: []byte("2")}}) != 2 {
		t.Fatal()
	}
}

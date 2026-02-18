package main

import (
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	ch := make(chan int, 1)
	go func() { ch <- Pool([]int{1, 2, 3}, 2) }()
	select {
	case v := <-ch:
		if v != 14 {
			t.Fatal()
		}
	case <-time.After(time.Second):
		t.Fatal("timeout")
	}
}

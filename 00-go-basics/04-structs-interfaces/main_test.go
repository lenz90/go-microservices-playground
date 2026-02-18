package main

import "testing"

func TestTotal(t *testing.T) {
	if Total([]Shape{Rect{2, 3}}) != 6 {
		t.Fatal()
	}
}

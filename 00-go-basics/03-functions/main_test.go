package main

import "testing"

func TestDivide(t *testing.T) {
	if v, _ := Divide(4, 2); v != 2 {
		t.Fatal()
	}
	if _, e := Divide(1, 0); e == nil {
		t.Fatal()
	}
}

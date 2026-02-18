package main

import "testing"

func TestSum(t *testing.T) {
	if Sum([]int{1, 2, 3}) != 6 {
		t.Fatal()
	}
}
func TestLookup(t *testing.T) {
	if v, ok := Lookup(map[string]int{"a": 1}, "a"); !ok || v != 1 {
		t.Fatal()
	}
}

package main

import "testing"

func TestRun(t *testing.T) {
	if s, _ := run([]string{"greet", "--name", "Ada"}); s != "hello, Ada" {
		t.Fatal()
	}
}

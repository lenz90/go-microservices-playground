package main

import "testing"

func TestEvent(t *testing.T) {
	if (Event{Type: "x", UserID: "u"}).UserID == "" {
		t.Fatal()
	}
}

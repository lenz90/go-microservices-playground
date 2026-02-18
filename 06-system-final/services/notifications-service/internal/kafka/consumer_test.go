package kafka

import "testing"

func TestExtractUserID(t *testing.T) {
	got := extractUserID([]byte(`{"user_id":"u1"}`))
	if got != "u1" {
		t.Fatalf("got %q", got)
	}
}

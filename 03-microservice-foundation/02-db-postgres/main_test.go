package main

import (
	"os"
	"testing"
)

func TestSkipIntegration(t *testing.T) {
	if os.Getenv("RUN_DB_TESTS") != "1" {
		t.Skip("set RUN_DB_TESTS=1")
	}
}

package main

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Setenv("PORT", "9999")
	os.Setenv("LOG_LEVEL", "debug")
	c, e := LoadConfig()
	if e != nil || c.Port != "9999" {
		t.Fatal()
	}
}

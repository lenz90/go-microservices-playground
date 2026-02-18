package main

import (
	"fmt"
	"log/slog"
	"os"
)

type Config struct {
	Port  string
	Level slog.Level
}

func LoadConfig() (Config, error) {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}
	lv := os.Getenv("LOG_LEVEL")
	if lv == "" {
		lv = "info"
	}
	var l slog.Level
	if err := l.UnmarshalText([]byte(lv)); err != nil {
		return Config{}, fmt.Errorf("invalid LOG_LEVEL: %w", err)
	}
	return Config{p, l}, nil
}
func main() { _, _ = LoadConfig() }

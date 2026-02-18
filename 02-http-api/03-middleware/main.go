package main

import (
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
)

func rec(l *slog.Logger) func(http.Handler) http.Handler {
	return func(n http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if x := recover(); x != nil {
					l.Error("panic", "x", x, "stack", string(debug.Stack()))
					http.Error(w, "internal", 500)
				}
			}()
			n.ServeHTTP(w, r)
		})
	}
}
func chain(h http.Handler, mw ...func(http.Handler) http.Handler) http.Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}
	return h
}
func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))
	_ = http.ListenAndServe(":8080", chain(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/panic" {
			panic("x")
		}
		w.Write([]byte("ok"))
	}), rec(l)))
}

package middleware

import "net/http"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-ID", "local")
		next.ServeHTTP(w, r)
	})
}

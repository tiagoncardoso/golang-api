package middleware

import (
	"golang.org/x/exp/slog"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request received", "method", r.Method, "url", r.URL.Path, "remote", r.RemoteAddr, "user-agent", r.UserAgent())
		next.ServeHTTP(w, r)
	})
}

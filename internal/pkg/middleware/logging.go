package middleware

import (
	"net/http"
	"time"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
)

func Logging(log logger.Logger) Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			crw := &CustomResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(crw, r)
			duration := time.Since(start)
			log.Info("HTTP Request",
				"http_method", r.Method,
				"http_path", r.URL.Path,
				"http_status", crw.statusCode,
				"http_latency_ms", duration.Milliseconds(),
			)
		})
	}
}

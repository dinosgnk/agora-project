package middleware

import (
	"net/http"
	"time"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
)

func HTTPLoggingMiddleware(log logger.Logger, path string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		responseWrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		handler(responseWrapper, r)
		duration := time.Since(start)

		log.Info("HTTP Request",
			"http_method", r.Method,
			"http_path", path,
			"http_status", responseWrapper.statusCode,
			"duration_ms", duration.Milliseconds(),
		)
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

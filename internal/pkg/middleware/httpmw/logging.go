package httpmw

import (
	"net/http"
	"time"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/middleware/logging"
)

func LoggingMiddleware(log logger.Logger) Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			crw := &CustomResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(crw, r)
			duration := time.Since(start)
			logging.LogRequest(log, r.Method, r.URL.Path, crw.statusCode, duration)
		})
	}
}

func TestMiddleware(log logger.Logger) Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info("Executing Test Middleware")
			next.ServeHTTP(w, r)
		})
	}
}

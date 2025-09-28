package httpmw

import (
	"net/http"
	"time"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/middleware"
	"github.com/dinosgnk/agora-project/internal/pkg/middleware/logging"
)

func LoggingMiddleware(log logger.Logger) middleware.Middleware {

	// Create the Middleware
	return func(next http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc the middleware returns
		return func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			responseWrapper := &CustomResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next(responseWrapper, r)
			duration := time.Since(start)

			logging.LogRequest(log, r.Method, r.URL.Path, responseWrapper.statusCode, duration)
		}
	}
}

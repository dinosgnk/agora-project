package logging

import (
	"time"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
)

func LogRequest(log logger.Logger, method, path string, statusCode int, duration time.Duration) {
	log.Info("HTTP Request",
		"http_method", method,
		"http_path", path,
		"http_status", statusCode,
		"http_latency_ms", duration.Milliseconds(),
	)
}

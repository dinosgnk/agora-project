package middleware

import (
	"time"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RequestLoggingMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		log.Info("HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency_ms", duration.Milliseconds(),
		)
	}
}

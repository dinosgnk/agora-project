package ginmw

import (
	"time"

	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/middleware/logging"
	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		logging.LogRequest(log, c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}

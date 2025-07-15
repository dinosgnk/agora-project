package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "catalog_http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path", "status_code"},
	)

	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "catalog_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status_code"},
	)

	httpResponseSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "catalog_http_response_size_bytes",
			Help:    "Size of HTTP responses in bytes",
			Buckets: []float64{100, 1000, 10000, 100000, 1000000, 10000000},
		},
		[]string{"method", "path", "status_code"},
	)

	httpRequestSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "catalog_http_request_size_bytes",
			Help:    "Size of HTTP requests in bytes",
			Buckets: []float64{100, 1000, 10000, 100000, 1000000},
		},
		[]string{"method", "path"},
	)

	httpRequestsInFlight = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "catalog_http_requests_in_flight",
			Help: "Number of HTTP requests currently being processed",
		},
		[]string{"method", "path"},
	)
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		method := ctx.Request.Method
		path := ctx.FullPath()

		if path == "" {
			path = "/"
		}

		requestSize := float64(ctx.Request.ContentLength)
		httpRequestSize.WithLabelValues(method, path).Observe(requestSize)

		httpRequestsInFlight.WithLabelValues(method, path).Inc()

		ctx.Next()

		httpRequestsInFlight.WithLabelValues(method, path).Dec()

		duration := time.Since(start).Seconds()
		statusCode := strconv.Itoa(ctx.Writer.Status())
		responseSize := float64(ctx.Writer.Size())

		httpRequestDuration.WithLabelValues(method, path, statusCode).Observe(duration)
		httpRequestsTotal.WithLabelValues(method, path, statusCode).Inc()
		httpResponseSize.WithLabelValues(method, path, statusCode).Observe(responseSize)
	}
}

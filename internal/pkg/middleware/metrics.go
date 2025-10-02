package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_milliseconds",
			Help:    "Duration of HTTP requests in milliseconds",
			Buckets: []float64{0.05, 0.1, 0.2, 0.5, 1, 2, 5, 10, 20, 50, 100},
		},
		[]string{"method", "path", "status_code", "service"},
	)

	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status_code", "service"},
	)

	httpResponseSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_size_bytes",
			Help:    "Size of HTTP responses in bytes",
			Buckets: []float64{100, 1000, 10000, 100000, 1000000, 10000000},
		},
		[]string{"method", "path", "status_code", "service"},
	)

	httpRequestSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_size_bytes",
			Help:    "Size of HTTP requests in bytes",
			Buckets: []float64{100, 1000, 10000, 100000, 1000000},
		},
		[]string{"method", "path", "status_code", "service"},
	)

	httpRequestsInFlight = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Number of HTTP requests currently being processed",
		},
		[]string{"method", "path", "service"},
	)
)

func Metrics(service string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			method := r.Method
			path := r.URL.Path

			requestSize := float64(r.ContentLength)

			crw := &CustomResponseWriter{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
				size:           0,
			}

			httpRequestsInFlight.WithLabelValues(method, path, service).Inc()
			next.ServeHTTP(crw, r)
			httpRequestsInFlight.WithLabelValues(method, path, service).Dec()

			duration := float64(time.Since(start).Nanoseconds()) / 1e6
			statusCode := strconv.Itoa(crw.statusCode/100) + "xx"
			responseSize := float64(crw.size)
			httpRequestDuration.WithLabelValues(method, path, statusCode, service).Observe(duration)
			httpRequestsTotal.WithLabelValues(method, path, statusCode, service).Inc()
			httpRequestSize.WithLabelValues(method, path, statusCode, service).Observe(requestSize)
			httpResponseSize.WithLabelValues(method, path, statusCode, service).Observe(responseSize)
		})
	}
}

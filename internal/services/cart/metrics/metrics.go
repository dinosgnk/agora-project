package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cart_http_request_duration_milliseconds",
			Help:    "Duration of HTTP requests in milliseconds",
			Buckets: []float64{0.05, 0.1, 0.2, 0.5, 1, 2, 5, 10, 20, 50, 100},
		},
		[]string{"method", "path", "status_code"},
	)

	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cart_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status_code"},
	)

	httpResponseSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cart_http_response_size_bytes",
			Help:    "Size of HTTP responses in bytes",
			Buckets: []float64{100, 1000, 10000, 100000, 1000000, 10000000},
		},
		[]string{"method", "path", "status_code"},
	)

	httpRequestSize = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "cart_http_request_size_bytes",
			Help:    "Size of HTTP requests in bytes",
			Buckets: []float64{100, 1000, 10000, 100000, 1000000},
		},
		[]string{"method", "path", "status_code"},
	)

	httpRequestsInFlight = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cart_http_requests_in_flight",
			Help: "Number of HTTP requests currently being processed",
		},
		[]string{"method", "path"},
	)
)

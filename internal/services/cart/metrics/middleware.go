package metrics

import (
	"net/http"
	"strconv"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *responseRecorder) Write(data []byte) (int, error) {
	size, err := r.ResponseWriter.Write(data)
	r.size += size
	return size, err
}

// HTTPMetricsMiddleware wraps HTTP handlers with metrics collection
func HTTPMetricsMiddleware(path string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := r.Method

		requestSize := float64(r.ContentLength)

		httpRequestsInFlight.WithLabelValues(method, path).Inc()
		defer httpRequestsInFlight.WithLabelValues(method, path).Dec()

		// Create a response recorder to capture status code and response size
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			size:           0,
		}

		next(recorder, r)

		duration := time.Since(start).Seconds()
		statusCode := strconv.Itoa(recorder.statusCode/100) + "xx"
		responseSize := float64(recorder.size)

		httpRequestDuration.WithLabelValues(method, path, statusCode).Observe(duration)
		httpRequestsTotal.WithLabelValues(method, path, statusCode).Inc()
		httpRequestSize.WithLabelValues(method, path, statusCode).Observe(requestSize)
		httpResponseSize.WithLabelValues(method, path, statusCode).Observe(responseSize)
	}
}

package metrics

import (
	"net/http"
	"strconv"
	"time"
)

const service string = "cart"

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (r *customResponseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *customResponseWriter) Write(data []byte) (int, error) {
	size, err := r.ResponseWriter.Write(data)
	r.size += size
	return size, err
}

func HTTPMetricsMiddleware(path string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := r.Method

		requestSize := float64(r.ContentLength)

		crw := &customResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			size:           0,
		}

		httpRequestsInFlight.WithLabelValues(method, path, service).Inc()
		next(crw, r)
		httpRequestsInFlight.WithLabelValues(method, path, service).Dec()

		duration := float64(time.Since(start).Nanoseconds()) / 1e6
		statusCode := strconv.Itoa(crw.statusCode/100) + "xx"
		responseSize := float64(crw.size)
		httpRequestDuration.WithLabelValues(method, path, statusCode, service).Observe(duration)
		httpRequestsTotal.WithLabelValues(method, path, statusCode, service).Inc()
		httpRequestSize.WithLabelValues(method, path, statusCode, service).Observe(requestSize)
		httpResponseSize.WithLabelValues(method, path, statusCode, service).Observe(responseSize)
	}
}

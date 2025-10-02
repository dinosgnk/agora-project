package middleware

import "net/http"

type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       float64
}

func (rw *CustomResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

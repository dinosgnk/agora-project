package httpmw

import "net/http"

type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *CustomResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

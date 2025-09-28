package httpx

import "net/http"

type ApiHandler interface {
	RegisterRoutes(mux *http.ServeMux) http.Handler
}

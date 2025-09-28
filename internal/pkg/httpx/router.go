package httpx

import (
	"net/http"

	"github.com/dinosgnk/agora-project/internal/pkg/middleware/httpmw"
)

type Router struct {
	mux         *http.ServeMux
	apiHandler  ApiHandler
	middlewares []httpmw.Middleware
}

func NewRouter(ah ApiHandler) *Router {
	mux := http.NewServeMux()

	return &Router{
		mux:         mux,
		apiHandler:  ah,
		middlewares: make([]httpmw.Middleware, 0),
	}
}

func (r *Router) Handle(pattern string, handler http.Handler) {
	r.mux.Handle(pattern, handler)
}

func (r *Router) AddMiddleware(middleware httpmw.Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) MiddlewareChain(next http.Handler, middlewares ...httpmw.Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}
	return next
}

func (r *Router) BuildHttpHandler() http.Handler {
	httpHandler := r.apiHandler.RegisterRoutes(r.mux)
	httpHandler = r.MiddlewareChain(httpHandler, r.middlewares...)
	return httpHandler
}

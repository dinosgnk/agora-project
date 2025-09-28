package server

import (
	"net/http"

	"github.com/dinosgnk/agora-project/internal/pkg/httpx"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/middleware/httpmw"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ServerConfig struct {
	address string
	port    string
}

type Server struct {
	httpHandler http.Handler
	log         logger.Logger
	apiHandler  httpx.ApiHandler
	ServerConfig
}

func NewServer(port string, apiHandler httpx.ApiHandler, log logger.Logger) *Server {
	router := httpx.NewRouter(apiHandler)
	router.Handle("/metrics", promhttp.Handler())
	router.AddMiddleware(httpmw.LoggingMiddleware(log))
	router.AddMiddleware(httpmw.TestMiddleware(log))

	httpHandler := router.BuildHttpHandler()

	return &Server{
		httpHandler: httpHandler,
		log:         log,
		apiHandler:  apiHandler,
		ServerConfig: ServerConfig{
			address: "localhost:" + port,
			port:    port,
		},
	}
}

func (s *Server) Run() error {
	s.log.Info("Starting server, listening on:", "address", s.address)
	if err := http.ListenAndServe(s.address, s.httpHandler); err != nil && err != http.ErrServerClosed {
		s.log.Error("Failed to start server", "error", err.Error())
		return err
	}
	return nil
}

// func (s *Server) Stop() {}

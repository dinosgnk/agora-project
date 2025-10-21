package server

import (
	"fmt"
	"net/http"

	"github.com/dinosgnk/agora-project/internal/pkg/httpx"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/dinosgnk/agora-project/internal/pkg/middleware"
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
	service     string
	ServerConfig
}

func NewServer(port string, apiHandler httpx.ApiHandler, log logger.Logger, service string) *Server {
	router := httpx.NewRouter(apiHandler)
	router.Handle("/metrics", promhttp.Handler())
	router.AddMiddleware(middleware.Logging(log))
	router.AddMiddleware(middleware.Metrics(service))

	httpHandler := router.BuildHttpHandler()

	return &Server{
		httpHandler: httpHandler,
		log:         log,
		apiHandler:  apiHandler,
		service:     service,
		ServerConfig: ServerConfig{
			address: "0.0.0.0:" + port,
			port:    port,
		},
	}
}

func (s *Server) Run() error {
	s.log.Info(fmt.Sprintf("Starting server, listening on: %s", s.address))
	if err := http.ListenAndServe(s.address, s.httpHandler); err != nil && err != http.ErrServerClosed {
		s.log.Error("Failed to start server", "error", err.Error())
		return err
	}
	return nil
}

// func (s *Server) Stop() {}

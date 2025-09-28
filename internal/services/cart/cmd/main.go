package main

import (
	"net/http"

	confighelper "github.com/dinosgnk/agora-project/internal/pkg/config"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	middleware "github.com/dinosgnk/agora-project/internal/pkg/middleware/logging"
	"github.com/dinosgnk/agora-project/internal/services/cart/config"
	"github.com/dinosgnk/agora-project/internal/services/cart/handler"
	"github.com/dinosgnk/agora-project/internal/services/cart/metrics"
	"github.com/dinosgnk/agora-project/internal/services/cart/repository"
	"github.com/dinosgnk/agora-project/internal/services/cart/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log := logger.NewLogger()

	cfg := confighelper.LoadConfig[config.AppConfig](log)

	log.Info("Starting cart service", "environment", cfg.Environment, "port", cfg.Port)

	cartRepository := repository.NewInMemoryRepository()
	cartService := service.NewCartService(cartRepository)
	cartHandler := handler.NewCartHandler(cartService, log)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/health", middleware.HTTPLoggingMiddleware(log, "/health", metrics.HTTPMetricsMiddleware("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Cart service is healthy"))
	})))

	http.HandleFunc("/cart", middleware.HTTPLoggingMiddleware(log, "/cart", metrics.HTTPMetricsMiddleware("/cart", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			cartHandler.GetCart(w, r)
		case http.MethodDelete:
			cartHandler.ClearCart(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	http.HandleFunc("/cart/item/add", middleware.HTTPLoggingMiddleware(log, "/cart/item/add", metrics.HTTPMetricsMiddleware("/cart/item/add", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			cartHandler.AddItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	http.HandleFunc("/cart/item/delete", middleware.HTTPLoggingMiddleware(log, "/cart/item/delete", metrics.HTTPMetricsMiddleware("/cart/item/delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			cartHandler.RemoveItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	log.Info("Cart service listening on port", "port", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Error("Failed to start server", "error", err.Error())
		panic(err)
	}
}

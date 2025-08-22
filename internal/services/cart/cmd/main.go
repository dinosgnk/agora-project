package main

import (
	"log"
	"net/http"

	"github.com/dinosgnk/agora-project/internal/services/cart/handler"
	"github.com/dinosgnk/agora-project/internal/services/cart/metrics"
	"github.com/dinosgnk/agora-project/internal/services/cart/repository"
	"github.com/dinosgnk/agora-project/internal/services/cart/service"
	"github.com/dinosgnk/agora-project/pkg/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := config.LoadConfig()
	log.Printf("Loaded config: %+v\n", cfg)

	cartRepository := repository.NewInMemoryRepository()
	cartService := service.NewCartService(cartRepository)
	cartHandler := handler.NewCartHandler(cartService)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/health", metrics.HTTPMetricsMiddleware("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Cart service is healthy"))
	}))

	http.HandleFunc("/cart", metrics.HTTPMetricsMiddleware("/cart", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			cartHandler.GetCart(w, r)
		case http.MethodDelete:
			cartHandler.ClearCart(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/cart/item/add", metrics.HTTPMetricsMiddleware("/cart/item/add", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			cartHandler.AddItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/cart/item/delete", metrics.HTTPMetricsMiddleware("/cart/item/delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			cartHandler.RemoveItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	log.Println("Cart service listening on port:", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}

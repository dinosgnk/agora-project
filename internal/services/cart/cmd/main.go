package main

import (
	"log"
	"net/http"

	"github.com/dinosgnk/agora-project/internal/services/cart/handler"
	"github.com/dinosgnk/agora-project/internal/services/cart/repository"
	"github.com/dinosgnk/agora-project/internal/services/cart/service"
)

func main() {
	cartRepository := repository.NewInMemoryRepository()
	cartService := service.NewCartService(cartRepository)
	cartHandler := handler.NewCartHandler(cartService)

	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			cartHandler.GetCart(w, r)
		case http.MethodDelete:
			cartHandler.ClearCart(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/cart/item/add", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			cartHandler.AddItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/cart/item/delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			cartHandler.RemoveItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Basket service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

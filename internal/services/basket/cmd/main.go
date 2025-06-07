package main

import (
	"log"
	"net/http"

	"agora/basket/handler"
	"agora/basket/repository"
	"agora/basket/service"
)

func main() {
	repo := repository.NewBasketRepository()
	svc := service.NewBasketService(repo)
	h := handler.NewBasketHandler(svc)

	http.HandleFunc("/basket", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetBasket(w, r)
		case http.MethodPost:
			h.AddItem(w, r)
		case http.MethodDelete:
			h.ClearBasket(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Basket service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package handlers

import (
	"Pet_Store/internal/repository"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	Repo repository.StoreRepository // Используем интерфейс, где есть товары
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Repo.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

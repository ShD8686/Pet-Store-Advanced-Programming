package handlers

import (
	"Pet_Store/internal/repository"
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	Repo repository.OrderRepository
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := h.Repo.GetAllOrders()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

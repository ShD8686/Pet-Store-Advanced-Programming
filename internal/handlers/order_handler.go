package handlers

import (
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"
	"encoding/json"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	Repo repository.OrderRepository
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := h.Repo.GetAllOrders()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) BuyProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		productID, _ := strconv.Atoi(r.FormValue("product_id"))

		// Создаем простой заказ (пока без авторизации, user_id = 1)
		newOrder := models.Order{
			PetID:  productID,
			UserID: 1,
			Total:  0, // Здесь можно вытянуть цену из базы
		}

		_ = h.Repo.CreateOrder(newOrder)

		// Перенаправляем на страницу "Спасибо за покупку" или обратно в маркет
		http.Redirect(w, r, "/view/products", http.StatusSeeOther)
	}
}

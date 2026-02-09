package handlers

import (
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"
	"html/template"
	"net/http"
)

type DashboardHandler struct {
	Repo *repository.SQLPetRepo
	Tmpl *template.Template
}

func (h *DashboardHandler) ViewDashboard(w http.ResponseWriter, r *http.Request) {
	// 1. Собираем данные
	orders, _ := h.Repo.GetUserOrders(1)
	appointments, _ := h.Repo.GetUserAppointments("User_1")

	// 2. Упаковываем всё в одну структуру для шаблона
	data := struct {
		Orders       []map[string]interface{}
		Appointments []models.Appointment
		UserName     string
	}{
		Orders:       orders,
		Appointments: appointments,
		UserName:     "User_1",
	}

	h.Tmpl.ExecuteTemplate(w, "dashboard.html", data)
}

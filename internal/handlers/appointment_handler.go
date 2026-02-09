package handlers

import (
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"
	"net/http"
)

type AppointmentHandler struct {
	Repo repository.StoreRepository
}

func (h *AppointmentHandler) BookAppointment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Читаем данные из формы
		app := models.Appointment{
			ServiceType:     r.FormValue("service_type"),
			PetName:         r.FormValue("pet_name"),
			OwnerName:       r.FormValue("owner_name"),
			AppointmentDate: r.FormValue("date"),
		}

		err := h.Repo.CreateAppointment(app)
		if err != nil {
			http.Error(w, "Ошибка при записи", http.StatusInternalServerError)
			return
		}

		// После успеха возвращаем на главную или страницу благодарности
		http.Redirect(w, r, "/view/pets", http.StatusSeeOther)
		return
	}
	// Если GET — просто показываем страницу (реализуем позже)
}

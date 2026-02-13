package handlers

import (
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"
	"html/template"
	"net/http"
)

type AppointmentHandler struct {
	Repo repository.StoreRepository
	Tmpl *template.Template
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

func (h *AppointmentHandler) ManageAppointments(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Получаем список всех записей из базы
		apps, _ := h.Repo.GetAllAppointments()
		// Отображаем страницу
		h.Tmpl.ExecuteTemplate(w, "appointments.html", apps)
		return
	}

	if r.Method == http.MethodPost {
		// Собираем данные из формы
		newApp := models.Appointment{
			ServiceType:     r.FormValue("service_type"),
			PetName:         r.FormValue("pet_name"),
			OwnerName:       r.FormValue("owner_name"),
			AppointmentDate: r.FormValue("date"),
			Status:          "pending",
		}

		// Сохраняем в базу
		_ = h.Repo.CreateAppointment(newApp)

		// Перенаправляем обратно на страницу записей, чтобы увидеть результат
		http.Redirect(w, r, "/view/appointments", http.StatusSeeOther)
	}
}

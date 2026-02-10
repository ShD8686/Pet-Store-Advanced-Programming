package handlers

import (
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"time"
)

type AuthHandler struct {
	Repo repository.PetRepository
	Tmpl *template.Template
}

// Register отвечает за регистрацию новых пользователей
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Если пользователь просто зашел на страницу /register
	if r.Method == http.MethodGet {
		err := h.Tmpl.ExecuteTemplate(w, "register.html", nil)
		if err != nil {
			http.Error(w, "Ошибка отображения страницы регистрации", http.StatusInternalServerError)
		}
		return
	}

	// Если пользователь отправил заполненную форму (кнопка "Создать аккаунт")
	if r.Method == http.MethodPost {
		u := models.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Role:     "user", // По умолчанию роль обычного пользователя
		}

		if err := h.Repo.CreateUser(u); err != nil {
			http.Error(w, "Пользователь с таким именем уже существует", http.StatusBadRequest)
			return
		}

		// После успешной регистрации отправляем на страницу входа
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// Login отвечает за вход в систему
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 1. ПРАВИЛЬНОЕ ОТОБРАЖЕНИЕ СТРАНИЦЫ (GET)
	if r.Method == http.MethodGet {
		err := h.Tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, "Ошибка отображения страницы входа", http.StatusInternalServerError)
		}
		return
	}

	// 2. ОБРАБОТКА ДАННЫХ ВХОДА (POST)
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Ищем пользователя в базе по логину
		user, err := h.Repo.GetUserByUsername(username)
		if err != nil {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		}

		// Сравниваем введенный пароль с зашифрованным паролем из базы
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
			return
		}

		// Если пароль подошел — создаем куку сессии (ID:Имя:Роль)
		cookie := &http.Cookie{
			Name:    "session",
			Value:   fmt.Sprintf("%d:%s:%s", user.ID, user.Username, user.Role),
			Expires: time.Now().Add(24 * time.Hour),
			Path:    "/",
		}
		http.SetCookie(w, cookie)

		// Перенаправляем на главную страницу с питомцами
		http.Redirect(w, r, "/view/pets", http.StatusSeeOther)
	}
}

// Logout отвечает за выход из системы (удаление куки)
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1, // Кука удалится мгновенно
		Path:   "/",
	})
	http.Redirect(w, r, "/view/pets", http.StatusSeeOther)
}
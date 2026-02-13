package handlers

import (
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Repo repository.PetRepository
}

func NewAuthHandler(repo repository.PetRepository) *AuthHandler {
	return &AuthHandler{Repo: repo}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user, err := h.Repo.GetUserByEmail(creds.Email)
	if err != nil {
		log.Printf("Login DB error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Сравниваем введенный пароль с хешем из базы данных
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Скрываем пароль перед отправкой на фронтенд
	user.Password = ""
	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Хешируем пароль перед сохранением
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error processing password", http.StatusInternalServerError)
		return
	}
	u.Password = string(hashedPassword)

	// Автоматическое определение админа по email
	if strings.Contains(strings.ToLower(u.Email), "admin") {
		u.Role = "admin"
	} else {
		u.Role = "user"
	}

	if err := h.Repo.CreateUser(u); err != nil {
		log.Printf("Registration error: %v", err)
		http.Error(w, "User already exists or DB error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created", "role": u.Role})
}

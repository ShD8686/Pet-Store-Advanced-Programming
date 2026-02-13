package handlers

import (
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"
	"encoding/json"
	"log"
	"net/http"
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

	if user == nil || user.Password != creds.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

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

	u.Role = "user"

	if err := h.Repo.CreateUser(u); err != nil {
		log.Printf("Registration error: %v", err) // Видим реальную причину в терминале
		http.Error(w, "User already exists or DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}

package handlers

import (
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"
	"encoding/json"
	"net/http"
)

type PetHandler struct {
	Repo repository.PetRepository
}

func (h *PetHandler) GetPets(w http.ResponseWriter, r *http.Request) {
	pets, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
}

func (h *PetHandler) CreatePet(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var p models.Pet
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if err := h.Repo.Create(p); err != nil {
        http.Error(w, "Failed to save pet", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Pet added successfully",
    })
}
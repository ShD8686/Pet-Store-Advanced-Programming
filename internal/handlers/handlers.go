package handlers

import (
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

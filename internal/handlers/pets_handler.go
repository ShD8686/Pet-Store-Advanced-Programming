package handlers

import (
	"Pet_Store/internal/repository"
	"encoding/json"
	"net/http"
)

type PetsHandler struct {
	Repo repository.PetRepository
}

func NewPetsHandler(repo repository.PetRepository) *PetsHandler {
	return &PetsHandler{Repo: repo}
}

func (h *PetsHandler) GetPets(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	pets, err := h.Repo.GetAll(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pets)
}

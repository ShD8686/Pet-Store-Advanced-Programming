package handlers

import (
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"
	"encoding/json"
	"net/http"
)

type ListingsHandler struct {
	Repo repository.PetRepository
}

func NewListingsHandler(repo repository.PetRepository) *ListingsHandler {
	return &ListingsHandler{Repo: repo}
}

func (h *ListingsHandler) GetListings(w http.ResponseWriter, r *http.Request) {
	listings, err := h.Repo.GetListings()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(listings)
}

func (h *ListingsHandler) CreateListing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var l models.Listing
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateListing(l); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Listing created successfully"})
}

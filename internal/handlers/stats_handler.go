package handlers

import (
	"Pet_Store/internal/repository"
	"encoding/json"
	"net/http"
)

type StatsHandler struct {
	Repo repository.PetRepository
}

func NewStatsHandler(repo repository.PetRepository) *StatsHandler {
	return &StatsHandler{Repo: repo}
}

func (h *StatsHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.Repo.GetStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stats)
}

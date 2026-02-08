package handlers

import (
    "encoding/json"
    "net/http"
)

type UserHandler struct {}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}
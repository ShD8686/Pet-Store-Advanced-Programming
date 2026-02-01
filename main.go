package main

import (
	"Pet_Store/internal/handlers"
	"Pet_Store/internal/repository"
	"fmt"
	"net/http"
)

func main() {
	repo := repository.NewMockPetRepo()
	petHandler := &handlers.PetHandler{Repo: repo}
	orderHandler := &handlers.OrderHandler{Repo: repo}

    http.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        petHandler.GetPets(w, r)
    } else if r.Method == http.MethodPost {
        petHandler.CreatePet(w, r)
    } else {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
    })

	http.HandleFunc("/orders", orderHandler.GetOrders)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Pet Store API! Use /pets to see the list.")
	})

	port := ":8080"
	fmt.Printf("Server is starting on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

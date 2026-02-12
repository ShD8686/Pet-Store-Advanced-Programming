package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"Pet_Store/internal/handlers"
	"Pet_Store/internal/repository"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./petstore.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	petRepo := repository.NewSQLitePetRepository(db)
	if err := petRepo.Seed(); err != nil {
		log.Printf("Error seeding data: %v", err)
	}

	petsHandler := handlers.NewPetsHandler(petRepo)
	statsHandler := handlers.NewStatsHandler(petRepo)
	listingsHandler := handlers.NewListingsHandler(petRepo)
	pageHandler := handlers.NewPageHandler()

	// API
	http.HandleFunc("/api/pets", handlers.CommonHeadersMiddleware(handlers.LoggerMiddleware(petsHandler.GetPets)))
	http.HandleFunc("/api/stats", handlers.CommonHeadersMiddleware(handlers.LoggerMiddleware(statsHandler.GetStats)))
	http.HandleFunc("/api/listings", handlers.CommonHeadersMiddleware(handlers.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			listingsHandler.CreateListing(w, r)
		} else {
			listingsHandler.GetListings(w, r)
		}
	})))

	// HTML Pages
	http.HandleFunc("/", pageHandler.IndexPage)
	http.HandleFunc("/info", pageHandler.InfoPage)
	http.HandleFunc("/stats", pageHandler.StatsPage)
	http.HandleFunc("/create-ad", pageHandler.CreateAdPage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("TAÑBA Server started on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

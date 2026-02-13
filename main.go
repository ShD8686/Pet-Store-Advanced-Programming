package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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

	// Инициализация таблиц
	if err := petRepo.InitSchema(); err != nil {
		log.Fatalf("Error initializing database schema: %v", err)
	}

	// Заполнение начальными данными
	if err := petRepo.Seed(); err != nil {
		log.Printf("Error seeding data: %v", err)
	}

	petsHandler := handlers.NewPetsHandler(petRepo)
	statsHandler := handlers.NewStatsHandler(petRepo)
	listingsHandler := handlers.NewListingsHandler(petRepo)
	authHandler := handlers.NewAuthHandler(petRepo)
	pageHandler := handlers.NewPageHandler()

	// API
	http.HandleFunc("/api/pets", handlers.CommonHeadersMiddleware(handlers.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			idStr := r.URL.Query().Get("id")
			id, _ := strconv.Atoi(idStr)
			petRepo.DeletePet(id)
			w.WriteHeader(http.StatusOK)
		} else {
			petsHandler.GetPets(w, r)
		}
	})))

	http.HandleFunc("/api/stats", handlers.CommonHeadersMiddleware(handlers.LoggerMiddleware(statsHandler.GetStats)))

	http.HandleFunc("/api/listings", handlers.CommonHeadersMiddleware(handlers.LoggerMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			listingsHandler.CreateListing(w, r)
		} else if r.Method == http.MethodDelete {
			idStr := r.URL.Query().Get("id")
			id, _ := strconv.Atoi(idStr)
			petRepo.DeleteListing(id)
			w.WriteHeader(http.StatusOK)
		} else {
			listingsHandler.GetListings(w, r)
		}
	})))

	http.HandleFunc("/api/login", handlers.CommonHeadersMiddleware(handlers.LoggerMiddleware(authHandler.Login)))
	http.HandleFunc("/api/register", handlers.CommonHeadersMiddleware(handlers.LoggerMiddleware(authHandler.Register)))

	// HTML Pages
	http.HandleFunc("/", pageHandler.IndexPage)
	http.HandleFunc("/info", pageHandler.InfoPage)
	http.HandleFunc("/stats", pageHandler.StatsPage)
	http.HandleFunc("/create-ad", pageHandler.CreateAdPage)
	http.HandleFunc("/login", pageHandler.LoginPage)
	http.HandleFunc("/register", pageHandler.RegisterPage)
	http.HandleFunc("/admin", pageHandler.AdminPage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("DNA Server started on port http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

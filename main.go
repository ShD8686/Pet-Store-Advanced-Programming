package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"Pet_Store/internal/handlers"
	"Pet_Store/internal/models"
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
	petRepo.InitSchema()
	petRepo.Seed()

	pageHandler := handlers.NewPageHandler()

	// API News
	http.HandleFunc("/api/news", handlers.CommonHeadersMiddleware(func(w http.ResponseWriter, r *http.Request) {
		news, _ := petRepo.GetNews()
		json.NewEncoder(w).Encode(news)
	}))

	// API Products
	http.HandleFunc("/api/products", handlers.CommonHeadersMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var p models.Product
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}
			if err := petRepo.AddProduct(p); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "Product added"})
		} else {
			cat := r.URL.Query().Get("category")
			products, _ := petRepo.GetProducts(cat)
			json.NewEncoder(w).Encode(products)
		}
	}))

	// API Appointments
	http.HandleFunc("/api/appointments", handlers.CommonHeadersMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var a models.Appointment
			json.NewDecoder(r.Body).Decode(&a)
			petRepo.CreateAppointment(a)
			w.WriteHeader(http.StatusCreated)
		} else {
			email := r.URL.Query().Get("email")
			apps, _ := petRepo.GetAppointmentsByEmail(email)
			json.NewEncoder(w).Encode(apps)
		}
	}))

	// Стандартные API
	http.HandleFunc("/api/pets", handlers.CommonHeadersMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			id, _ := strconv.Atoi(r.URL.Query().Get("id"))
			petRepo.DeletePet(id)
		} else {
			status := r.URL.Query().Get("status")
			pets, _ := petRepo.GetAll(status)
			json.NewEncoder(w).Encode(pets)
		}
	}))
	http.HandleFunc("/api/stats", handlers.CommonHeadersMiddleware(func(w http.ResponseWriter, r *http.Request) {
		stats, _ := petRepo.GetStats()
		json.NewEncoder(w).Encode(stats)
	}))
	http.HandleFunc("/api/listings", handlers.CommonHeadersMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var l models.Listing
			json.NewDecoder(r.Body).Decode(&l)
			petRepo.CreateListing(l)
		} else if r.Method == http.MethodDelete {
			id, _ := strconv.Atoi(r.URL.Query().Get("id"))
			petRepo.DeleteListing(id)
		} else {
			listings, _ := petRepo.GetListings()
			json.NewEncoder(w).Encode(listings)
		}
	}))
	http.HandleFunc("/api/login", handlers.CommonHeadersMiddleware(handlers.NewAuthHandler(petRepo).Login))
	http.HandleFunc("/api/register", handlers.CommonHeadersMiddleware(handlers.NewAuthHandler(petRepo).Register))

	// Pages
	http.HandleFunc("/", pageHandler.IndexPage)
	http.HandleFunc("/info", pageHandler.InfoPage)
	http.HandleFunc("/stats", pageHandler.StatsPage)
	http.HandleFunc("/create-ad", pageHandler.CreateAdPage)
	http.HandleFunc("/login", pageHandler.LoginPage)
	http.HandleFunc("/register", pageHandler.RegisterPage)
	http.HandleFunc("/admin", pageHandler.AdminPage)

	http.HandleFunc("/vet", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/vet.html")
	})
	http.HandleFunc("/shop", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/templates/shop.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("DNA Server running on http://localhost:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

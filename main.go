package main

import (
	"Pet_Store/internal/handlers"
	"Pet_Store/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	
	db, err := sql.Open("sqlite3", "./pet_store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	
	schema := `
	CREATE TABLE IF NOT EXISTS pets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		category TEXT,
		price REAL,
		status TEXT
	);
	`
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	
	petRepo := repository.NewSQLitePetRepo(db)

	petHandler := &handlers.PetHandler{Repo: petRepo}
	orderHandler := &handlers.OrderHandler{Repo: petRepo}
	userHandler := &handlers.UserHandler{}

	
	go func() {
		for {
			time.Sleep(30 * time.Second)
			fmt.Println("[Background Worker] System is running OK")
		}
	}()

	
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
	http.HandleFunc("/register", userHandler.Register)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


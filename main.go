package main

import (
	"Pet_Store/internal/handlers"
	"Pet_Store/internal/repository"
	"database/sql"
	"fmt"
	"html/template"
	_ "html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

func main() {
	db, err := sql.Open("sqlite", "./pet_store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Читаем схему из файла
	schema, err := os.ReadFile("sql/schema.sql")
	if err == nil {
		db.Exec(string(schema))
		fmt.Println("Database schema updated from sql/schema.sql")
	}

	// Инициализация ОДНОГО репозитория для всех
	sqlRepo := repository.NewSQLPetRepo(db)

	// Передаем один и тот же sqlRepo в разные хендлеры
	petHandler := &handlers.PetHandler{Repo: sqlRepo}
	orderHandler := &handlers.OrderHandler{Repo: sqlRepo}

	// В main.go, там где создаются хендлеры:
	productHandler := &handlers.ProductHandler{Repo: sqlRepo}
	appointmentHandler := &handlers.AppointmentHandler{Repo: sqlRepo}

	go func() {
		for {
			fmt.Println("[System]: Background check OK.")
			time.Sleep(1 * time.Minute)
		}
	}()

	http.HandleFunc("/view/pets", func(w http.ResponseWriter, r *http.Request) {
		pets, err := sqlRepo.GetAllPets() // Берем данные из базы
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Рендерим HTML и передаем туда список pets
		tmpl.ExecuteTemplate(w, "index.html", pets)
	})

	http.HandleFunc("/view/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := sqlRepo.GetAllProducts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "products.html", products)
	})

	http.HandleFunc("/book", appointmentHandler.BookAppointment)
	http.HandleFunc("/pets", petHandler.GetPets)
	http.HandleFunc("/orders", orderHandler.GetOrders)
	http.HandleFunc("/register", (&handlers.UserHandler{}).Register)
	// В блоке эндпоинтов:
	http.HandleFunc("/products", productHandler.GetProducts)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pet Store API is running!")
	})

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

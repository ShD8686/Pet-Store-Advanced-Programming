package main

import (
	"Pet_Store/internal/handlers"
	"Pet_Store/internal/repository"
	"database/sql"
	"fmt"
	"html/template"
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

	schema, err := os.ReadFile("sql/schema.sql")
	if err == nil {
		db.Exec(string(schema))
		fmt.Println("Database schema updated from sql/schema.sql")
	}


	sqlRepo := repository.NewSQLPetRepo(db)


	petHandler := &handlers.PetHandler{Repo: sqlRepo}
	orderHandler := &handlers.OrderHandler{Repo: sqlRepo}


	productHandler := &handlers.ProductHandler{Repo: sqlRepo}
	appHandler := &handlers.AppointmentHandler{Repo: sqlRepo, Tmpl: tmpl}
	dashHandler := &handlers.DashboardHandler{Repo: sqlRepo, Tmpl: tmpl}

	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	go func() {
		for {
			fmt.Println("[System]: Background check OK.")
			time.Sleep(1 * time.Minute)
		}
	}()

	http.HandleFunc("/view/pets", func(w http.ResponseWriter, r *http.Request) {
		pets, err := sqlRepo.GetAllPets()
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

	// Один роут для просмотра и для создания (GET и POST)
	http.HandleFunc("/view/appointments", appHandler.ManageAppointments)
	http.HandleFunc("/book", appHandler.ManageAppointments)

	http.HandleFunc("/view/shelter", func(w http.ResponseWriter, r *http.Request) {
		pets, _ := sqlRepo.GetPetsForAdoption()
		// Используем тот же index.html, но передаем только животных из приюта
		tmpl.ExecuteTemplate(w, "index.html", pets)
	})

	http.HandleFunc("/buy", orderHandler.BuyProduct)
	http.HandleFunc("/view/dashboard", dashHandler.ViewDashboard)

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// 1. Получаем текст поиска из URL (например, /search?q=food)
		query := r.URL.Query().Get("q")

		// 2. Вызываем метод поиска из репозитория
		// Используем sqlRepo, который у тебя уже инициализирован выше в main
		foundProducts, err := sqlRepo.SearchProducts(query)
		if err != nil {
			http.Error(w, "Ошибка поиска", http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(w, "products.html", foundProducts)
	})

	http.HandleFunc("/pets", petHandler.GetPets)
	http.HandleFunc("/orders", orderHandler.GetOrders)
	http.HandleFunc("/register", (&handlers.UserHandler{}).Register)

	http.HandleFunc("/products", productHandler.GetProducts)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pet Store API is running!")
	})

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

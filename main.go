package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"log"
	"net/http"
	"os"
	"strconv"

	"Pet_Store/internal/handlers"
	"Pet_Store/internal/models"
	"Pet_Store/internal/repository"

	_ "modernc.org/sqlite" // Используем современный драйвер SQLite
	_ "modernc.org/sqlite"
)

// 1. Глобальная переменная для шаблонов (ищет все .html в папке web/templates)
var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

// --- Вспомогательные функции для авторизации и защиты ---

// GetCurrentUser вытягивает имя и роль из куки "session"
func GetCurrentUser(r *http.Request) (string, string) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", ""
	}
	// Ожидаемый формат куки: "id:username:role"
	parts := strings.Split(cookie.Value, ":")
	if len(parts) < 3 {
		return "", ""
	}
	return parts[1], parts[2] // Возвращаем username и role
}

// AdminOnly — защита для страниц, доступных только админу
func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, role := GetCurrentUser(r)
		if role != "admin" {
			http.Error(w, "Доступ запрещен: только для администраторов", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

// --- ОСНОВНАЯ ФУНКЦИЯ ---

func main() {
	// 1. Подключение к базе данных
	db, err := sql.Open("sqlite", "./pet_store.db")
func main() {
	db, err := sql.Open("sqlite", "./petstore.db")
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}
	defer db.Close()

	// 2. Инициализация таблиц из файла schema.sql (если файл существует)
	schemaFile, err := os.ReadFile("sql/schema.sql")
	if err == nil {
		_, err = db.Exec(string(schemaFile))
		if err != nil {
			log.Println("Заметка по SQL: ", err)
		} else {
			fmt.Println("База данных успешно обновлена (schema.sql).")
		}
	}

	// 3. Создание репозитория
	sqlRepo := repository.NewSQLPetRepo(db)

	// 4. ИНИЦИАЛИЗАЦИЯ ХЕНДЛЕРОВ (Передаем Repo и Tmpl)
	authHandler := &handlers.AuthHandler{
		Repo: sqlRepo, 
		Tmpl: tmpl, // ПЕРЕДАЕМ ШАБЛОНЫ, чтобы работали Register и Login
	}
	petHandler := &handlers.PetHandler{Repo: sqlRepo}
	orderHandler := &handlers.OrderHandler{Repo: sqlRepo}
	productHandler := &handlers.ProductHandler{Repo: sqlRepo}
	appHandler := &handlers.AppointmentHandler{Repo: sqlRepo, Tmpl: tmpl}
	dashHandler := &handlers.DashboardHandler{Repo: sqlRepo, Tmpl: tmpl}

	// 5. ФОНОВАЯ ЗАДАЧА (Goroutine)
	go func() {
		for {
			fmt.Println("[System]: Background Check OK -", time.Now().Format("15:04:05"))
			time.Sleep(1 * time.Minute)
		}
	}()

	// 6. РОУТИНГ (МАРШРУТЫ)

	// --- Авторизация (Логин, Регистрация, Выход) ---
	http.HandleFunc("/login", authHandler.Login)
	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/logout", authHandler.Logout)

	// --- Визуальные страницы (Frontend) ---
	http.HandleFunc("/view/pets", func(w http.ResponseWriter, r *http.Request) {
		pets, _ := sqlRepo.GetAllPets()
		username, role := GetCurrentUser(r)
		
		// Передаем в HTML данные о пользователе и его роли
		data := map[string]interface{}{
			"Pets":     pets,
			"Username": username,
			"IsAdmin":  role == "admin",
		}
		tmpl.ExecuteTemplate(w, "index.html", data)
	})

	// Пример админской страницы
	http.HandleFunc("/admin/dashboard", AdminOnly(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Это панель администратора. Вы видите это, потому что вы — Админ.")
	}))

	// --- API эндпоинты (JSON данные) ---
	http.HandleFunc("/view/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := sqlRepo.GetAllProducts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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

	http.HandleFunc("/pets", petHandler.GetPets)
	http.HandleFunc("/products", productHandler.GetProducts)
	http.HandleFunc("/orders", orderHandler.GetOrders)

	// Перенаправление с главной на список питомцев
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/view/pets", http.StatusSeeOther)
	})

	// 7. ЗАПУСК СЕРВЕРА
	fmt.Println("Сервер работает на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
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

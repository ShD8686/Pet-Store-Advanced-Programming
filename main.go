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
	"strings"
	"time"

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
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}
	defer db.Close()

	schema, err := os.ReadFile("sql/schema.sql")
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

	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

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
	http.HandleFunc("/register", (&handlers.UserHandler{}).Register)

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
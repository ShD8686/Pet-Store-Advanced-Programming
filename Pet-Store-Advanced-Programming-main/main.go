package main

import (
	"Pet_Store/internal/handlers"
	"Pet_Store/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

func main() {
	// 1. Инициализация базы данных
	db, err := sql.Open("sqlite", "./pet_store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создаем таблицу, если её нет (Persistence)
	createTable := `CREATE TABLE IF NOT EXISTS pets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		category TEXT,
		price REAL,
		status TEXT
	);`
	db.Exec(createTable)

	// 2. Инициализация репозитория и хендлеров
	// Теперь используем SQL репозиторий вместо Mock
	petRepo := repository.NewSQLPetRepo(db)
	petHandler := &handlers.PetHandler{Repo: petRepo}

	// Для заказов пока можно оставить Mock или тоже переписать на SQL
	orderRepo := repository.NewMockPetRepo()
	orderHandler := &handlers.OrderHandler{Repo: orderRepo}

	// 3. Concurrency (Requirement 3: Goroutine)
	go func() {
		for {
			fmt.Println("[Background Task]: Updating pet inventory statistics...")
			time.Sleep(30 * time.Second) // Выполняется каждые 30 секунд
		}
	}()

	// 4. Эндпоинты
	http.HandleFunc("/pets", petHandler.GetPets)
	http.HandleFunc("/orders", orderHandler.GetOrders)
	http.HandleFunc("/register", (&handlers.UserHandler{}).Register)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Pet Store API!")
	})

	port := ":8080"
	fmt.Printf("Server is starting on http://localhost%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

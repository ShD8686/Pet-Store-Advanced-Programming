package repository

import (
	"Pet_Store/internal/models"
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type SQLitePetRepository struct {
	DB *sql.DB
}

func NewSQLitePetRepository(db *sql.DB) *SQLitePetRepository {
	return &SQLitePetRepository{DB: db}
}

func (r *SQLitePetRepository) InitSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS pets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		chip_number TEXT UNIQUE,
		name TEXT,
		type TEXT,
		gender TEXT,
		breed TEXT,
		status TEXT,
		image_url TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS listings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type TEXT NOT NULL,
		pet_type TEXT NOT NULL,
		breed TEXT,
		photo_url TEXT NOT NULL,
		reward REAL DEFAULT 0,
		price REAL DEFAULT 0,
		has_insurance BOOLEAN DEFAULT 0,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE,
		password TEXT,
		role TEXT DEFAULT 'user',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		category TEXT NOT NULL,
		price REAL NOT NULL,
		image_url TEXT,
		description TEXT
	);

	CREATE TABLE IF NOT EXISTS appointments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_email TEXT NOT NULL,
		pet_name TEXT,
		vet_name TEXT,
		date TEXT,
		time TEXT,
		reason TEXT,
		status TEXT DEFAULT 'pending'
	);

	CREATE TABLE IF NOT EXISTS news (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		category TEXT,
		date TEXT,
		image_url TEXT,
		summary TEXT
	);`

	_, err := r.DB.Exec(schema)
	return err
}

func (r *SQLitePetRepository) GetNews() ([]models.News, error) {
	rows, err := r.DB.Query("SELECT id, title, category, date, image_url, summary FROM news ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var news []models.News
	for rows.Next() {
		var n models.News
		rows.Scan(&n.ID, &n.Title, &n.Category, &n.Date, &n.ImageURL, &n.Summary)
		news = append(news, n)
	}
	return news, nil
}

func (r *SQLitePetRepository) GetProducts(category string) ([]models.Product, error) {
	query := "SELECT id, name, category, price, image_url, description FROM products"
	var rows *sql.Rows
	var err error
	if category != "" {
		query += " WHERE category = ?"
		rows, err = r.DB.Query(query, category)
	} else {
		rows, err = r.DB.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.ImageURL, &p.Description)
		products = append(products, p)
	}
	return products, nil
}

func (r *SQLitePetRepository) AddProduct(p models.Product) error {
	_, err := r.DB.Exec(`INSERT INTO products (name, category, price, image_url, description) VALUES (?, ?, ?, ?, ?)`,
		p.Name, p.Category, p.Price, p.ImageURL, p.Description)
	return err
}

func (r *SQLitePetRepository) CreateAppointment(a models.Appointment) error {
	_, err := r.DB.Exec(`INSERT INTO appointments (user_email, pet_name, vet_name, date, time, reason) VALUES (?, ?, ?, ?, ?, ?)`,
		a.UserEmail, a.PetName, a.VetName, a.Date, a.Time, a.Reason)
	return err
}

func (r *SQLitePetRepository) GetAppointmentsByEmail(email string) ([]models.Appointment, error) {
	rows, err := r.DB.Query("SELECT id, user_email, pet_name, vet_name, date, time, reason, status FROM appointments WHERE user_email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []models.Appointment
	for rows.Next() {
		var a models.Appointment
		rows.Scan(&a.ID, &a.UserEmail, &a.PetName, &a.VetName, &a.Date, &a.Time, &a.Reason, &a.Status)
		apps = append(apps, a)
	}
	return apps, nil
}

func (r *SQLitePetRepository) GetAll(status string) ([]models.Pet, error) {
	query := "SELECT id, chip_number, name, type, gender, breed, status, image_url FROM pets"
	var rows *sql.Rows
	var err error
	if status != "" {
		query += " WHERE status = ?"
		rows, err = r.DB.Query(query, status)
	} else {
		rows, err = r.DB.Query(query)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var pets []models.Pet
	for rows.Next() {
		var p models.Pet
		rows.Scan(&p.ID, &p.ChipNumber, &p.Name, &p.Type, &p.Gender, &p.Breed, &p.Status, &p.ImageURL)
		pets = append(pets, p)
	}
	return pets, nil
}

func (r *SQLitePetRepository) GetStats() (models.Stats, error) {
	var total, cats, dogs int
	r.DB.QueryRow("SELECT COUNT(*) FROM pets").Scan(&total)
	r.DB.QueryRow("SELECT COUNT(*) FROM pets WHERE type = 'кошка'").Scan(&cats)
	r.DB.QueryRow("SELECT COUNT(*) FROM pets WHERE type = 'собака'").Scan(&dogs)
	return models.Stats{TotalRegistered: total, TotalCats: cats, TotalDogs: dogs, LastUpdate: time.Now().Format(time.RFC3339)}, nil
}

func (r *SQLitePetRepository) CreateListing(l models.Listing) error {
	_, err := r.DB.Exec(`INSERT INTO listings (type, pet_type, breed, photo_url, reward, price, has_insurance, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		l.Type, l.PetType, l.Breed, l.PhotoURL, l.Reward, l.Price, l.HasInsurance, l.Description)
	return err
}

func (r *SQLitePetRepository) GetListings() ([]models.Listing, error) {
	rows, err := r.DB.Query("SELECT id, type, pet_type, breed, photo_url, reward, price, has_insurance, description, created_at FROM listings ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listings []models.Listing
	for rows.Next() {
		var l models.Listing
		rows.Scan(&l.ID, &l.Type, &l.PetType, &l.Breed, &l.PhotoURL, &l.Reward, &l.Price, &l.HasInsurance, &l.Description, &l.CreatedAt)
		listings = append(listings, l)
	}
	return listings, nil
}

func (r *SQLitePetRepository) DeleteListing(id int) error {
	_, err := r.DB.Exec("DELETE FROM listings WHERE id = ?", id)
	return err
}

func (r *SQLitePetRepository) CreateUser(u models.User) error {
	if u.Role == "" {
		u.Role = "user"
	}
	_, err := r.DB.Exec("INSERT INTO users (email, password, role) VALUES (?, ?, ?)", u.Email, u.Password, u.Role)
	return err
}

func (r *SQLitePetRepository) GetUserByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.DB.QueryRow("SELECT id, email, password, role FROM users WHERE email = ?", email).Scan(&u.ID, &u.Email, &u.Password, &u.Role)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *SQLitePetRepository) DeletePet(id int) error {
	_, err := r.DB.Exec("DELETE FROM pets WHERE id = ?", id)
	return err
}

func (r *SQLitePetRepository) AddPet(p models.Pet) error {
	_, err := r.DB.Exec("INSERT INTO pets (chip_number, name, type, gender, breed, status, image_url) VALUES (?, ?, ?, ?, ?, ?, ?)",
		p.ChipNumber, p.Name, p.Type, p.Gender, p.Breed, p.Status, p.ImageURL)
	return err
}

func (r *SQLitePetRepository) Seed() error {
	var count int
	r.DB.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	if count == 0 {
		prods := []models.Product{
			{Name: "Корм Royal Canin (2кг)", Category: "food", Price: 12500, ImageURL: "https://images.unsplash.com/photo-1589924691995-400dc9ecc119?auto=format&fit=crop&w=400&q=80", Description: "Сбалансированный корм для взрослых собак"},
			{Name: "Витамины Beaphar", Category: "medicine", Price: 4500, ImageURL: "https://images.unsplash.com/photo-1584308666744-24d5c474f2ae?auto=format&fit=crop&w=400&q=80", Description: "Комплекс витаминов для кошек"},
			{Name: "Игрушка-пищалка", Category: "toy", Price: 2200, ImageURL: "https://images.unsplash.com/photo-1576201836106-db1758fd1c97?auto=format&fit=crop&w=400&q=80", Description: "Прочная резиновая игрушка"},
		}
		for _, p := range prods {
			r.DB.Exec("INSERT INTO products (name, category, price, image_url, description) VALUES (?, ?, ?, ?, ?)", p.Name, p.Category, p.Price, p.ImageURL, p.Description)
		}
	}

	r.DB.QueryRow("SELECT COUNT(*) FROM news").Scan(&count)
	if count == 0 {
		newsItems := []models.News{
			{Title: "Всеобщая вакцинация в Астане", Category: "События", Date: "25 Янв", ImageURL: "https://images.unsplash.com/photo-1628033033904-91496a793f77?auto=format&fit=crop&w=800&q=80", Summary: "Запущена программа бесплатной вакцинации против бешенства."},
			{Title: "Как выбрать правильный корм?", Category: "Советы", Date: "12 Янв", ImageURL: "https://images.unsplash.com/photo-1583337130417-3346a1be7dee?auto=format&fit=crop&w=800&q=80", Summary: "Гид по выбору питания для щенков и взрослых собак."},
			{Title: "Открытие новой клиники в Алматы", Category: "Новости", Date: "05 Янв", ImageURL: "https://images.unsplash.com/photo-1516733725897-1aa73b87c8e8?auto=format&fit=crop&w=800&q=80", Summary: "Современный центр ветеринарии открыл свои двери."},
		}
		for _, n := range newsItems {
			r.DB.Exec("INSERT INTO news (title, category, date, image_url, summary) VALUES (?, ?, ?, ?, ?)", n.Title, n.Category, n.Date, n.ImageURL, n.Summary)
		}
	}

	r.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if count == 0 {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		r.CreateUser(models.User{Email: "admin@tanba.kz", Password: string(hashed), Role: "admin"})
	}
	return nil
}

package repository

import (
	"Pet_Store/internal/models"
	"database/sql"
	"time"
)

type SQLitePetRepository struct {
	DB *sql.DB
}

func NewSQLitePetRepository(db *sql.DB) *SQLitePetRepository {
	return &SQLitePetRepository{DB: db}
}

// InitSchema создает таблицы, если они не существуют
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
	);`

	_, err := r.DB.Exec(schema)
	return err
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

	return models.Stats{
		TotalRegistered: total,
		TotalCats:       cats,
		TotalDogs:       dogs,
		LastUpdate:      time.Now().Format(time.RFC3339),
	}, nil
}

func (r *SQLitePetRepository) CreateListing(l models.Listing) error {
	_, err := r.DB.Exec(`
		INSERT INTO listings (type, pet_type, breed, photo_url, reward, price, has_insurance, description)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
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
	// Проверяем наличие пользователей
	err := r.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// Default admin
		r.CreateUser(models.User{Email: "admin@tanba.kz", Password: "admin", Role: "admin"})
		r.CreateUser(models.User{Email: "user@test.kz", Password: "user", Role: "user"})
	}

	r.DB.QueryRow("SELECT COUNT(*) FROM pets").Scan(&count)
	if count > 0 {
		return nil
	}

	seeds := []models.Pet{
		{ChipNumber: "398010001", Name: "Луна", Type: "кошка", Gender: "female", Breed: "Бенгальская", Status: "lost", ImageURL: "https://images.unsplash.com/photo-1513245543132-31f507417b26?auto=format&fit=crop&w=500&q=80"},
		{ChipNumber: "398010002", Name: "Арчи", Type: "собака", Gender: "male", Breed: "Золотистый ретривер", Status: "found", ImageURL: "https://images.unsplash.com/photo-1552053831-71594a27632d?auto=format&fit=crop&w=500&q=80"},
	}

	for _, p := range seeds {
		r.AddPet(p)
	}
	return nil
}

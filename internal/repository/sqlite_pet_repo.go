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

func (r *SQLitePetRepository) Seed() error {
	var count int
	r.DB.QueryRow("SELECT COUNT(*) FROM pets").Scan(&count)
	if count > 0 {
		return nil
	}

	seeds := []models.Pet{
		{ChipNumber: "398010001", Name: "Луна", Type: "кошка", Gender: "female", Breed: "Бенгальская", Status: "lost", ImageURL: "https://images.unsplash.com/photo-1513245543132-31f507417b26?auto=format&fit=crop&w=500&q=80"},
		{ChipNumber: "398010002", Name: "Арчи", Type: "собака", Gender: "male", Breed: "Золотистый ретривер", Status: "found", ImageURL: "https://images.unsplash.com/photo-1552053831-71594a27632d?auto=format&fit=crop&w=500&q=80"},
	}

	for _, p := range seeds {
		_, err := r.DB.Exec("INSERT INTO pets (chip_number, name, type, gender, breed, status, image_url) VALUES (?, ?, ?, ?, ?, ?, ?)",
			p.ChipNumber, p.Name, p.Type, p.Gender, p.Breed, p.Status, p.ImageURL)
		if err != nil {
			return err
		}
	}
	return nil
}

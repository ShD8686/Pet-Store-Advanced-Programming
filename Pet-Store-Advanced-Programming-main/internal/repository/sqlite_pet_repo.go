package repository

import (
	"Pet_Store/internal/models"
	"database/sql"
)

type SQLPetRepo struct {
	DB *sql.DB
}

func NewSQLPetRepo(db *sql.DB) *SQLPetRepo {
	return &SQLPetRepo{DB: db}
}

func (r *SQLPetRepo) GetAll() ([]models.Pet, error) {
	rows, err := r.DB.Query("SELECT id, name, category, price, status FROM pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Status); err != nil {
			return nil, err
		}
		pets = append(pets, p)
	}
	return pets, nil
}

func (r *SQLPetRepo) GetByID(id int) (models.Pet, error) {
	var p models.Pet
	err := r.DB.QueryRow("SELECT id, name, category, price, status FROM pets WHERE id = ?", id).
		Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Status)
	return p, err
}

func (r *SQLPetRepo) Create(pet models.Pet) error {
	_, err := r.DB.Exec("INSERT INTO pets (name, category, price, status) VALUES (?, ?, ?, ?)",
		pet.Name, pet.Category, pet.Price, pet.Status)
	return err
}

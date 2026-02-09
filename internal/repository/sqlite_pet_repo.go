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

func (r *SQLPetRepo) GetAllPets() ([]models.Pet, error) {
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

func (r *SQLPetRepo) CreatePet(pet models.Pet) error {
	_, err := r.DB.Exec("INSERT INTO pets (name, category, price, status) VALUES (?, ?, ?, ?)",
		pet.Name, pet.Category, pet.Price, pet.Status)
	return err
}

func (r *SQLPetRepo) GetPetsForAdoption() ([]models.Pet, error) {
	return nil, nil
}

// Реализация для OrderRepository
func (r *SQLPetRepo) GetAllOrders() ([]models.Order, error) {
	return nil, nil
}

func (r *SQLPetRepo) CreateOrder(order models.Order) error {
	return nil
}

func (r *SQLPetRepo) GetAllProducts() ([]models.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, category, price, stock, description FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Stock, &p.Description); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// Остальные заглушки для StoreRepository
func (r *SQLPetRepo) CreateProduct(product models.Product) error { return nil }

func (r *SQLPetRepo) CreateAppointment(app models.Appointment) error {
	_, err := r.DB.Exec(`
		INSERT INTO appointments (service_type, pet_name, owner_name, appointment_date, status) 
		VALUES (?, ?, ?, ?, ?)`,
		app.ServiceType, app.PetName, app.OwnerName, app.AppointmentDate, "pending")
	return err
}

func (r *SQLPetRepo) GetAllAppointments() ([]models.Appointment, error) { return nil, nil }

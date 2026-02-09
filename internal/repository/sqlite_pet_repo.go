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
	// Фильтруем по полю is_for_adoption
	rows, err := r.DB.Query("SELECT id, name, category, price, status, description FROM pets WHERE is_for_adoption = 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []models.Pet
	for rows.Next() {
		var p models.Pet
		rows.Scan(&p.ID, &p.Name, &p.Category, &p.Price, &p.Status, &p.Description)
		p.IsForAdoption = true
		pets = append(pets, p)
	}
	return pets, nil
}

// Реализация для OrderRepository
func (r *SQLPetRepo) GetAllOrders() ([]models.Order, error) {
	return nil, nil
}

func (r *SQLPetRepo) CreateOrder(order models.Order) error {
	_, err := r.DB.Exec("INSERT INTO orders (pet_id, user_id, total_price) VALUES (?, ?, ?)",
		order.PetID, order.UserID, order.Total)
	return err
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

func (r *SQLPetRepo) GetAllAppointments() ([]models.Appointment, error) {
	rows, err := r.DB.Query("SELECT id, service_type, pet_name, owner_name, appointment_date, status FROM appointments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []models.Appointment
	for rows.Next() {
		var a models.Appointment
		rows.Scan(&a.ID, &a.ServiceType, &a.PetName, &a.OwnerName, &a.AppointmentDate, &a.Status)
		apps = append(apps, a)
	}
	return apps, nil
}

// Получить историю заказов с названиями товаров
func (r *SQLPetRepo) GetUserOrders(userID int) ([]map[string]interface{}, error) {
	query := `
		SELECT o.id, p.name, o.total_price 
		FROM orders o
		JOIN products p ON o.pet_id = p.id
		WHERE o.user_id = ?`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		var price float64
		rows.Scan(&id, &name, &price)
		results = append(results, map[string]interface{}{
			"ID": id, "ProductName": name, "Price": price,
		})
	}
	return results, nil
}

// Получить записи к врачу по имени владельца (пока у нас нет системы логина по ID)
func (r *SQLPetRepo) GetUserAppointments(ownerName string) ([]models.Appointment, error) {
	rows, err := r.DB.Query("SELECT service_type, pet_name, appointment_date, status FROM appointments WHERE owner_name = ?", ownerName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []models.Appointment
	for rows.Next() {
		var a models.Appointment
		rows.Scan(&a.ServiceType, &a.PetName, &a.AppointmentDate, &a.Status)
		apps = append(apps, a)
	}
	return apps, nil
}

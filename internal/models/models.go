package models

type Pet struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	Status        string  `json:"status"`
	Description   string  `json:"description"`
	IsForAdoption bool    `json:"is_for_adoption"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

type Appointment struct {
	ID              int    `json:"id"`
	ServiceType     string `json:"service_type"`
	PetName         string `json:"pet_name"`
	OwnerName       string `json:"owner_name"`
	AppointmentDate string `json:"appointment_date"`
	Status          string `json:"status"`
}

// Добавили структуру Order, чтобы ошибки в Handler исчезли
type Order struct {
	ID     int     `json:"id"`
	PetID  int     `json:"pet_id"`
	UserID int     `json:"user_id"`
	Total  float64 `json:"total_price"`
}

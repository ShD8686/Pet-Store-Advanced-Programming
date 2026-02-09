package repository

import "Pet_Store/internal/models"

type PetRepository interface {
	GetAllPets() ([]models.Pet, error)
	GetPetsForAdoption() ([]models.Pet, error)
	CreatePet(pet models.Pet) error
}

type OrderRepository interface {
	GetAllOrders() ([]models.Order, error)
	CreateOrder(order models.Order) error
}

// Дополнительный интерфейс для товаров и записей (на будущее)
type StoreRepository interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(product models.Product) error
	CreateAppointment(app models.Appointment) error
	GetAllAppointments() ([]models.Appointment, error)

	SearchProducts(query string) ([]models.Product, error)
}

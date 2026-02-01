package repository

import "Pet_Store/internal/models"

type PetRepository interface {
	GetAll() ([]models.Pet, error)
	GetByID(id int) (models.Pet, error)
	Create(pet models.Pet) error
}

type OrderRepository interface {
	GetAllOrders() ([]models.Order, error)
}

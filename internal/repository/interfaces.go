package repository

import "Pet_Store/internal/models"

type PetRepository interface {
	GetAll() ([]models.Pet, error)
	GetByID(id int) (models.Pet, error)
	Create(pet models.Pet) error
}
<<<<<<< HEAD
=======
type OrderRepository interface {
	GetAllOrders() ([]models.Order, error)
}
>>>>>>> a734a48007c034570d1b6d56cc88d3f837dcaad3

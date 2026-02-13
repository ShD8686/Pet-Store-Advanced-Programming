package repository

import "Pet_Store/internal/models"

type PetRepository interface {
	GetAll(status string) ([]models.Pet, error)
	GetStats() (models.Stats, error)
	Seed() error

	CreateListing(l models.Listing) error
	GetListings() ([]models.Listing, error)
	DeleteListing(id int) error

	CreateUser(u models.User) error
	GetUserByEmail(email string) (*models.User, error)
	DeletePet(id int) error
	AddPet(p models.Pet) error

	GetProducts(category string) ([]models.Product, error)
	AddProduct(p models.Product) error
	CreateAppointment(a models.Appointment) error
	GetAppointmentsByEmail(email string) ([]models.Appointment, error)
}

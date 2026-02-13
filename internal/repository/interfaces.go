package repository

import "Pet_Store/internal/models"

type PetRepository interface {
	GetAll(status string) ([]models.Pet, error)
	GetStats() (models.Stats, error)
	Seed() error

	// Listings methods
	CreateListing(l models.Listing) error
	GetListings() ([]models.Listing, error)
	DeleteListing(id int) error

	// User methods
	CreateUser(u models.User) error
	GetUserByEmail(email string) (*models.User, error)
	DeletePet(id int) error
	AddPet(p models.Pet) error
}

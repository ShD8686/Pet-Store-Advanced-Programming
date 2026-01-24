package repository

import "Pet_Store/internal/models"

type CategoryRepository interface {
	GetAll() ([]string, error)
	GetByName(name string) (*models.Pet, error)
}

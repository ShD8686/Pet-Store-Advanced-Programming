package repository

import "Pet_Store/internal/models"

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetByID(id int) (*models.Category, error)
}

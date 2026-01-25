package repository

import (
	"Pet_Store/internal/models"
	"fmt"
)

type MockPetRepo struct {
	pets []models.Pet
}

func NewMockPetRepo() *MockPetRepo {
	return &MockPetRepo{
		pets: []models.Pet{
			{ID: 1, Name: "Pet1", Category: "Parrot", Price: 100.50, Status: "Available"},
			{ID: 2, Name: "Pet2", Category: "Cat", Price: 50.00, Status: "Available"},
			{ID: 3, Name: "Pet3", Category: "Dog", Price: 75.00, Status: "Available"},
		},
	}
}

func (m *MockPetRepo) GetAll() ([]models.Pet, error) {
	return m.pets, nil
}

func (m *MockPetRepo) GetByID(id int) (models.Pet, error) {
	for _, p := range m.pets {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Pet{}, fmt.Errorf("pet with id %d not found", id)
}

func (m *MockPetRepo) Create(pet models.Pet) error {
	m.pets = append(m.pets, pet)
	return nil
}

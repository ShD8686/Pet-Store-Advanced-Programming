package models

import "time"

type Pet struct {
	ID         int    `json:"id"`
	ChipNumber string `json:"chip_number"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Gender     string `json:"gender"`
	Breed      string `json:"breed"`
	Status     string `json:"status"` // 'lost', 'found', 'registered'
	ImageURL   string `json:"image_url"`
}

type Listing struct {
	ID           int       `json:"id"`
	Type         string    `json:"type"`          // 'lost', 'sell', 'give', 'found'
	PetType      string    `json:"pet_type"`      // 'dog', 'cat', etc.
	Breed        string    `json:"breed"`         // optional
	PhotoURL     string    `json:"photo_url"`     // required
	Reward       float64   `json:"reward"`        // optional (for 'lost')
	Price        float64   `json:"price"`         // optional (for 'sell')
	HasInsurance bool      `json:"has_insurance"` // optional (for 'sell', 'give')
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
}

type Stats struct {
	TotalRegistered int    `json:"total"`
	TotalCats       int    `json:"cats"`
	TotalDogs       int    `json:"dogs"`
	LastUpdate      string `json:"last_update"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"` // Изменено с "-" на "password"
	Role     string `json:"role"`
}

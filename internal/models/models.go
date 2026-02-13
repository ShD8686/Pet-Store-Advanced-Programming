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
	Type         string    `json:"type"`
	PetType      string    `json:"pet_type"`
	Breed        string    `json:"breed"`
	PhotoURL     string    `json:"photo_url"`
	Reward       float64   `json:"reward"`
	Price        float64   `json:"price"`
	HasInsurance bool      `json:"has_insurance"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
}

type Appointment struct {
	ID        int    `json:"id"`
	UserEmail string `json:"user_email"`
	PetName   string `json:"pet_name"`
	VetName   string `json:"vet_name"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Reason    string `json:"reason"`
	Status    string `json:"status"` // 'pending', 'confirmed'
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"` // 'food', 'medicine', 'toy'
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	Description string  `json:"description"`
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
	Password string `json:"password"`
	Role     string `json:"role"`
}

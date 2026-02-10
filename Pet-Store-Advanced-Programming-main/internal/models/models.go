package models

type Pet struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Status   string  `json:"status"`
}

type Order struct {
	ID     int     `json:"id"`
	PetID  int     `json:"pet_id"`
	UserID int     `json:"user_id"`
	Total  float64 `json:"total_price"`
}

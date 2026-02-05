package models

import "time"

type Product struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId *int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	Category   *Category  `json:"category"`
}

type Products []Product

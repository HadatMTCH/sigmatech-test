package models

import "time"

type Product struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Category  string    `db:"category" json:"category"`
	MinPrice  float64   `db:"min_price" json:"min_price"`
	MaxPrice  float64   `db:"max_price" json:"max_price"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

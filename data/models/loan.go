package models

import "time"

type LoanApplication struct {
	ID              int       `db:"id" json:"id"`
	UserID          int       `db:"user_id" json:"user_id"`
	ProductID       int       `db:"product_id" json:"product_id"`
	Amount          float64   `db:"amount" json:"amount"`
	Tenor           int       `db:"tenor" json:"tenor"`
	Status          string    `db:"status" json:"status"`
	ApplicationDate time.Time `db:"application_date" json:"application_date"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type CreateLoanApplicationRequest struct {
	ProductID int     `json:"product_id"`
	Amount    float64 `json:"amount"`
	Tenor     int     `json:"tenor"`
}

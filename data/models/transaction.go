package models

import "time"

type Transaction struct {
	ID                int       `db:"id" json:"id"`
	UserID            int       `db:"user_id" json:"user_id"`
	LoanApplicationID int       `db:"loan_application_id" json:"loan_application_id"`
	Amount            float64   `db:"amount" json:"amount"`
	Type              string    `db:"type" json:"type"`
	TransactionDate   time.Time `db:"transaction_date" json:"transaction_date"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

type CreateTransactionRequest struct {
	LoanApplicationID int     `json:"loan_application_id"`
	Amount            float64 `json:"amount"`
	Type              string  `json:"type"`
}

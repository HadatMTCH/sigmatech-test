package models

import "time"

type CreditLimit struct {
	ID          int       `db:"id" json:"id"`
	UserID      int       `db:"user_id" json:"user_id"`
	Limit1Month float64   `db:"limit_1_month" json:"limit_1_month"`
	Limit2Month float64   `db:"limit_2_month" json:"limit_2_month"`
	Limit3Month float64   `db:"limit_3_month" json:"limit_3_month"`
	Limit6Month float64   `db:"limit_6_month" json:"limit_6_month"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

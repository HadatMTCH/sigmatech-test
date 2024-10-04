package models

import (
	"time"
)

type User struct {
	ID           int       `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	PasswordHash string    `db:"password_hash" json:"-"` // Don't expose password hash in JSON
	FullName     string    `db:"full_name" json:"full_name"`
	NIK          string    `db:"nik" json:"nik"`
	BirthDate    time.Time `db:"birth_date" json:"birth_date"`
	Salary       float64   `db:"salary" json:"salary"`
	KTPPhoto     []byte    `db:"ktp_photo" json:"-"`    // Don't expose in JSON, handle separately
	SelfiePhoto  []byte    `db:"selfie_photo" json:"-"` // Don't expose in JSON, handle separately
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type JWTClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type JWTRequest struct {
	ID       int
	Username string
	FullName string
}

package template

import (
	"base-api/data/models"
	"base-api/infra/db"
	"context"
)

type Template interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateLoanApplication(ctx context.Context, app models.LoanApplication) (int, error)
	GetLoanApplications(ctx context.Context, userID int) ([]models.LoanApplication, error)
	GetProducts(ctx context.Context) ([]models.Product, error)
	CreateTransaction(ctx context.Context, transaction models.Transaction) (int, error)
	GetTransactions(ctx context.Context, userID int) ([]models.Transaction, error)
	GetCreditLimit(ctx context.Context, userID int) (*models.CreditLimit, error)
	UpdateCreditLimit(ctx context.Context, limit models.CreditLimit) error
}

func New(db *db.DB) Template {
	return &template{
		db,
	}
}

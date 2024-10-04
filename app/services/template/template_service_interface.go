package template

import (
	"base-api/data/models"
	"base-api/infra/context/repository"
	"context"
)

type Template interface {
	RegisterUser(ctx context.Context, user models.User) (int, error)
	AuthenticateUser(ctx context.Context, username, password string) (*models.User, error)
	CreateLoanApplication(ctx context.Context, userID int, app models.CreateLoanApplicationRequest) (int, error)
	GetLoanApplications(ctx context.Context, userID int) ([]models.LoanApplication, error)
	GetProducts(ctx context.Context) ([]models.Product, error)
	CreateTransaction(ctx context.Context, userID int, req models.CreateTransactionRequest) (int, error)
	GetTransactions(ctx context.Context, userID int) ([]models.Transaction, error)
	GetCreditLimit(ctx context.Context, userID int) (*models.CreditLimit, error)
	UpdateCreditLimit(ctx context.Context, limit models.CreditLimit) error
}

func New(ctx *repository.RepositoryContext) Template {
	return &template{
		ctx,
	}
}

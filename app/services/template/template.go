package template

import (
	"base-api/data/models"
	"base-api/infra/context/repository"
	"base-api/utils"
	"context"
	"errors"
	"time"
)

type template struct {
	*repository.RepositoryContext
}

func (s *template) RegisterUser(ctx context.Context, user models.User) (int, error) {
	hashedPassword, err := utils.HashPassword(user.PasswordHash)
	if err != nil {
		return 0, err
	}
	user.PasswordHash = hashedPassword
	return s.RepositoryContext.TemplateRepository.CreateUser(ctx, user)
}

func (s *template) AuthenticateUser(ctx context.Context, username, password string) (*models.User, error) {
	user, err := s.RepositoryContext.TemplateRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *template) CreateLoanApplication(ctx context.Context, userID int, app models.CreateLoanApplicationRequest) (int, error) {
	// Check credit limit
	limit, err := s.RepositoryContext.TemplateRepository.GetCreditLimit(ctx, userID)
	if err != nil {
		return 0, err
	}

	// Simplified credit limit check (you might want to make this more sophisticated)
	var allowedLimit float64
	switch app.Tenor {
	case 1:
		allowedLimit = limit.Limit1Month
	case 2:
		allowedLimit = limit.Limit2Month
	case 3:
		allowedLimit = limit.Limit3Month
	case 6:
		allowedLimit = limit.Limit6Month
	default:
		return 0, errors.New("invalid tenor")
	}

	if app.Amount > allowedLimit {
		return 0, errors.New("loan amount exceeds credit limit")
	}

	loanApp := models.LoanApplication{
		UserID:          userID,
		ProductID:       app.ProductID,
		Amount:          app.Amount,
		Tenor:           app.Tenor,
		Status:          "PENDING",
		ApplicationDate: time.Now(),
	}

	return s.RepositoryContext.TemplateRepository.CreateLoanApplication(ctx, loanApp)
}

func (s *template) GetLoanApplications(ctx context.Context, userID int) ([]models.LoanApplication, error) {
	return s.RepositoryContext.TemplateRepository.GetLoanApplications(ctx, userID)
}

func (s *template) GetProducts(ctx context.Context) ([]models.Product, error) {
	return s.RepositoryContext.TemplateRepository.GetProducts(ctx)
}

func (s *template) CreateTransaction(ctx context.Context, userID int, req models.CreateTransactionRequest) (int, error) {
	// You might want to add more validation here, e.g., checking if the loan application exists and belongs to the user
	transaction := models.Transaction{
		UserID:            userID,
		LoanApplicationID: req.LoanApplicationID,
		Amount:            req.Amount,
		Type:              req.Type,
		TransactionDate:   time.Now(),
	}
	return s.RepositoryContext.TemplateRepository.CreateTransaction(ctx, transaction)
}

func (s *template) GetTransactions(ctx context.Context, userID int) ([]models.Transaction, error) {
	return s.RepositoryContext.TemplateRepository.GetTransactions(ctx, userID)
}

func (s *template) GetCreditLimit(ctx context.Context, userID int) (*models.CreditLimit, error) {
	return s.RepositoryContext.TemplateRepository.GetCreditLimit(ctx, userID)
}

func (s *template) UpdateCreditLimit(ctx context.Context, limit models.CreditLimit) error {
	return s.RepositoryContext.TemplateRepository.UpdateCreditLimit(ctx, limit)
}

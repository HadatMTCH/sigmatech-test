package template

import (
	"base-api/data/models"
	"base-api/infra/db"
	"context"
	"database/sql"
	"errors"
	"time"
)

type template struct {
	*db.DB
}

func (r *template) CreateUser(ctx context.Context, user models.User) (int, error) {
	query := `INSERT INTO users (username, password_hash, full_name, nik, birth_date, salary, ktp_photo, selfie_photo, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $9) RETURNING id`
	var userID int
	err := r.DB.QueryRowContext(ctx, query, user.Username, user.PasswordHash, user.FullName, user.NIK, user.BirthDate, user.Salary, user.KTPPhoto, user.SelfiePhoto, time.Now()).Scan(&userID)
	return userID, err
}

func (r *template) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE username = $1 LIMIT 1`
	err := r.DB.Master().GetContext(ctx, &user, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *template) CreateLoanApplication(ctx context.Context, app models.LoanApplication) (int, error) {
	query := `INSERT INTO loan_applications (user_id, product_id, amount, tenor, status, application_date, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $7) RETURNING id`
	var appID int
	err := r.DB.QueryRowContext(ctx, query, app.UserID, app.ProductID, app.Amount, app.Tenor, app.Status, app.ApplicationDate, time.Now()).Scan(&appID)
	return appID, err
}

func (r *template) GetLoanApplications(ctx context.Context, userID int) ([]models.LoanApplication, error) {
	var apps []models.LoanApplication
	query := `SELECT * FROM loan_applications WHERE user_id = $1`
	err := r.DB.Master().SelectContext(ctx, &apps, query, userID)
	return apps, err
}

func (r *template) GetProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	query := `SELECT * FROM products`
	err := r.DB.Master().SelectContext(ctx, &products, query)
	return products, err
}

func (r *template) CreateTransaction(ctx context.Context, transaction models.Transaction) (int, error) {
	query := `INSERT INTO transactions (user_id, loan_application_id, amount, type, transaction_date, created_at, updated_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $6) RETURNING id`
	var transactionID int
	err := r.DB.QueryRowContext(ctx, query, transaction.UserID, transaction.LoanApplicationID, transaction.Amount, transaction.Type, transaction.TransactionDate, time.Now()).Scan(&transactionID)
	return transactionID, err
}

func (r *template) GetTransactions(ctx context.Context, userID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := `SELECT * FROM transactions WHERE user_id = $1`
	err := r.DB.Master().SelectContext(ctx, &transactions, query, userID)
	return transactions, err
}

func (r *template) GetCreditLimit(ctx context.Context, userID int) (*models.CreditLimit, error) {
	var limit models.CreditLimit
	query := `SELECT * FROM credit_limits WHERE user_id = $1 LIMIT 1`
	err := r.DB.Master().GetContext(ctx, &limit, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("credit limit not found")
		}
		return nil, err
	}
	return &limit, nil
}

func (r *template) UpdateCreditLimit(ctx context.Context, limit models.CreditLimit) error {
	query := `UPDATE credit_limits SET
	          limit_1_month = $1, limit_2_month = $2, limit_3_month = $3, limit_6_month = $4, updated_at = $5
	          WHERE user_id = $6`
	_, err := r.DB.ExecContext(ctx, query, limit.Limit1Month, limit.Limit2Month, limit.Limit3Month, limit.Limit6Month, time.Now(), limit.UserID)
	return err
}

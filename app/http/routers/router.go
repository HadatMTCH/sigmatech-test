package routers

import (
	infra "base-api/infra/context"

	"github.com/gorilla/mux"
)

const (
	GET   = "GET"
	POST  = "POST"
	PUT   = "PUT"
	PATCH = "PATCH"
)

// InitialRouter for object routers
func InitialRouter(infra infra.InfraContextInterface, r *mux.Router) *mux.Router {
	s := r.PathPrefix("/api").Subrouter()

	// Auth routes
	auth := s.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", infra.Handler().TemplateHandler.RegisterUser).Methods(POST)
	auth.HandleFunc("/login", infra.Handler().TemplateHandler.Login).Methods(POST)

	// Protected routes
	protected := s.PathPrefix("/loan").Subrouter()
	protected.Use(infra.Middleware().TokenMiddleware.TokenAuthorize)

	// Loan application routes
	protected.HandleFunc("/applications", infra.Handler().TemplateHandler.CreateLoanApplication).Methods(POST)
	protected.HandleFunc("/applications", infra.Handler().TemplateHandler.GetLoanApplications).Methods(GET)

	// Product routes
	protected.HandleFunc("/products", infra.Handler().TemplateHandler.GetProducts).Methods(GET)

	// Transaction routes
	protected.HandleFunc("/transactions", infra.Handler().TemplateHandler.CreateTransaction).Methods(POST)
	protected.HandleFunc("/transactions", infra.Handler().TemplateHandler.GetTransactions).Methods(GET)

	// Credit limit routes
	protected.HandleFunc("/credit-limit", infra.Handler().TemplateHandler.GetCreditLimit).Methods(GET)
	protected.HandleFunc("/credit-limit", infra.Handler().TemplateHandler.UpdateCreditLimit).Methods(PUT)

	return r
}

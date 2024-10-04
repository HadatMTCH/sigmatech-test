package template

import (
	"base-api/infra/context/service"
	"net/http"
)

type Template interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	CreateLoanApplication(w http.ResponseWriter, r *http.Request)
	GetLoanApplications(w http.ResponseWriter, r *http.Request)
	GetProducts(w http.ResponseWriter, r *http.Request)
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactions(w http.ResponseWriter, r *http.Request)
	GetCreditLimit(w http.ResponseWriter, r *http.Request)
	UpdateCreditLimit(w http.ResponseWriter, r *http.Request)
}

func New(serviceCtx *service.ServiceContext) Template {
	return &template{
		serviceCtx,
	}
}

package template

import (
	"base-api/data/models"
	"base-api/infra/context/service"
	"base-api/infra/middleware"
	"base-api/utils"
	"encoding/json"
	"net/http"
)

type template struct {
	*service.ServiceContext
}

func (h *template) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		res := utils.Response{
			Code: http.StatusBadRequest,
			Data: nil,
			Err:  utils.STATUS_BAD_REQUEST,
			Msg:  "Invalid request payload",
		}
		res.JSONResponse(w)
		return
	}

	userID, err := h.ServiceContext.TemplateService.RegisterUser(r.Context(), user)
	if err != nil {
		res := utils.Response{
			Code: http.StatusInternalServerError,
			Data: nil,
			Err:  utils.STATUS_INTERNAL_ERR,
			Msg:  err.Error(),
		}
		res.JSONResponse(w)
		return
	}

	res := utils.Response{
		Code: http.StatusCreated,
		Data: map[string]int{"user_id": userID},
		Msg:  "User registered successfully",
	}
	res.JSONResponse(w)
}

func (h *template) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		res := utils.Response{
			Code: http.StatusBadRequest,
			Data: nil,
			Err:  utils.STATUS_BAD_REQUEST,
			Msg:  "Invalid request payload",
		}
		res.JSONResponse(w)
		return
	}

	user, err := h.ServiceContext.TemplateService.AuthenticateUser(r.Context(), loginReq.Username, loginReq.Password)
	if err != nil {
		res := utils.Response{
			Code: http.StatusUnauthorized,
			Data: nil,
			Err:  utils.STATUS_UNAUTHORIZED,
			Msg:  "Invalid credentials",
		}
		res.JSONResponse(w)
		return
	}

	token, err := h.JWTService.GenerateJWTToken(r.Context(), models.JWTRequest{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
	})
	if err != nil {
		res := utils.Response{
			Code: http.StatusInternalServerError,
			Data: nil,
			Err:  utils.STATUS_INTERNAL_ERR,
			Msg:  "Error generating token",
		}
		res.JSONResponse(w)
		return
	}

	res := utils.Response{
		Code: http.StatusOK,
		Data: models.LoginResponse{
			ID:    user.ID,
			Token: token,
		},
		Msg: "Login successful",
	}
	res.JSONResponse(w)
}

func (h *template) CreateLoanApplication(w http.ResponseWriter, r *http.Request) {
	token := middleware.GetTokenFromContext(r.Context())
	var appReq models.CreateLoanApplicationRequest
	err := json.NewDecoder(r.Body).Decode(&appReq)
	if err != nil {
		res := utils.Response{
			Code: http.StatusBadRequest,
			Data: nil,
			Err:  utils.STATUS_BAD_REQUEST,
			Msg:  "Invalid request payload",
		}
		res.JSONResponse(w)
		return
	}

	appID, err := h.ServiceContext.TemplateService.CreateLoanApplication(r.Context(), token.ID, appReq)
	if err != nil {
		res := utils.Response{
			Code: http.StatusInternalServerError,
			Data: nil,
			Err:  utils.STATUS_INTERNAL_ERR,
			Msg:  err.Error(),
		}
		res.JSONResponse(w)
		return
	}

	res := utils.Response{
		Code: http.StatusCreated,
		Data: map[string]int{"application_id": appID},
		Msg:  "Loan application created successfully",
	}
	res.JSONResponse(w)
}

func (h *template) GetLoanApplications(w http.ResponseWriter, r *http.Request) {
	token := middleware.GetTokenFromContext(r.Context())
	apps, err := h.ServiceContext.TemplateService.GetLoanApplications(r.Context(), token.ID)
	if err != nil {
		res := utils.Response{
			Code: http.StatusInternalServerError,
			Data: nil,
			Err:  utils.STATUS_INTERNAL_ERR,
			Msg:  err.Error(),
		}
		res.JSONResponse(w)
		return
	}

	res := utils.Response{
		Code: http.StatusOK,
		Data: apps,
		Msg:  "Loan applications retrieved successfully",
	}
	res.JSONResponse(w)
}

func (h *template) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ServiceContext.TemplateService.GetProducts(r.Context())
	if err != nil {
		res := utils.Response{
			Code: http.StatusInternalServerError,
			Data: nil,
			Err:  utils.STATUS_INTERNAL_ERR,
			Msg:  err.Error(),
		}
		res.JSONResponse(w)
		return
	}

	res := utils.Response{
		Code: http.StatusOK,
		Data: products,
		Msg:  "Products retrieved successfully",
	}
	res.JSONResponse(w)
}

func (h *template) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	token := middleware.GetTokenFromContext(r.Context())
	var transReq models.CreateTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transReq)
	if err != nil {
		res := utils.Response{
			Code: http.StatusBadRequest,
			Data: nil,
			Err:  utils.STATUS_BAD_REQUEST,
			Msg:  "Invalid request payload",
		}
		res.JSONResponse(w)
		return
	}

	transID, err := h.ServiceContext.TemplateService.CreateTransaction(r.Context(), token.ID, transReq)
	if err != nil {
		res := utils.Response{
			Code: http.StatusInternalServerError,
			Data: nil,
			Err:  utils.STATUS_INTERNAL_ERR,
			Msg:  err.Error(),
		}
		res.JSONResponse(w)
		return
	}

	res := utils.Response{
		Code: http.StatusCreated,
		Data: map[string]int{"transaction_id": transID},
		Msg:  "Transaction created successfully",
	}
	res.JSONResponse(w)
}

func (h *template) GetTransactions(w http.ResponseWriter, r *http.Request) {
	token := middleware.GetTokenFromContext(r.Context())
	transactions, err := h.ServiceContext.TemplateService.GetTransactions(r.Context(), token.ID)
	if err != nil {
		res := utils.Response{
			Code: http.StatusInternalServerError,
			Data: nil,
			Err:  utils.STATUS_INTERNAL_ERR,
			Msg:  err.Error(),
		}
		res.JSONResponse(w)
		return
	}

	res := utils.Response{
		Code: http.StatusOK,
		Data: transactions,
		Msg:  "Transactions retrieved successfully",
	}
	res.JSONResponse(w)
}

func (h *template) GetCreditLimit(w http.ResponseWriter, r *http.Request) {
	token := middleware.GetTokenFromContext(r.Context())
	limit, err := h.ServiceContext.TemplateService.GetCreditLimit(r.Context(), token.ID)
	if err != nil {
		res := utils.Response{
			Code: http.StatusInternalServerError,
			Data: nil,
			Err:  utils.STATUS_INTERNAL_ERR,
			Msg:  err.Error(),
		}
		res.JSONResponse(w)
		return
	}

	res := utils.Response{
		Code: http.StatusOK,
		Data: limit,
		Msg:  "Credit limit retrieved successfully",
	}
	res.JSONResponse(w)
}

func (h *template) UpdateCreditLimit(w http.ResponseWriter, r *http.Request) {
	token := middleware.GetTokenFromContext(r.Context())
	var creditLimit models.CreditLimit
	err := json.NewDecoder(r.Body).Decode(&creditLimit)
	if err != nil {
		res := utils.Response{
			Code: http.StatusBadRequest,
			Data: nil,
			Err:  utils.STATUS_BAD_REQUEST,
			Msg:  "Invalid request payload",
		}
		res.JSONResponse(w)
		return
	}

	// Ensure the user can only update their own credit limit
	creditLimit.UserID = token.ID

	err = h.ServiceContext.TemplateService.UpdateCreditLimit(r.Context(), creditLimit)
	if err != nil {
		res := utils.Response{
			Code: http.StatusInternalServerError,
			Data: nil,
			Err:  utils.STATUS_INTERNAL_ERR,
			Msg:  err.Error(),
		}
		res.JSONResponse(w)
		return
	}

	res := utils.Response{
		Code: http.StatusOK,
		Data: nil,
		Msg:  "Credit limit updated successfully",
	}
	res.JSONResponse(w)
}

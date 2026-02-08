package handler

import (
	"basic-go-api/internal/model"
	"basic-go-api/internal/service"
	"encoding/json"
	"net/http"
)

type CheckoutRequest struct {
	Items []model.CheckoutItem `json:"items"`
}

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (h *TransactionHandler) Transactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		model.JSONResponse(w, http.StatusMethodNotAllowed, false, "method not allowed", nil)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	ctx := r.Context()
	transaction, err := h.service.Create(ctx, req.Items)
	if err != nil {
		model.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusCreated, true, "transaction created", transaction)
}

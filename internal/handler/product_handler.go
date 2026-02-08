package handler

import (
	"basic-go-api/internal/model"
	"basic-go-api/internal/service"
	"encoding/json"
	"net/http"
	"strings"
)

type ProductRequest struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	Stock       int      `json:"stock"`
	CategoryIDs []string `json:"category_ids"`
}

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) Products(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAll(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		model.JSONResponse(w, http.StatusMethodNotAllowed, false, "method not allowed", nil)
	}
}

func (h *ProductHandler) ProductByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")
	if id == "" {
		model.JSONResponse(w, http.StatusNotFound, false, "not found", nil)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getByID(w, r, id)
	case http.MethodPut:
		h.update(w, r, id)
	case http.MethodDelete:
		h.delete(w, r, id)
	default:
		model.JSONResponse(w, http.StatusMethodNotAllowed, false, "method not allowed", nil)
	}
}

func (h *ProductHandler) getAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := r.URL.Query().Get("name")

	data, err := h.service.GetAll(ctx, name)
	if err != nil {
		model.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusOK, true, "success", data)
}

func (h *ProductHandler) create(w http.ResponseWriter, r *http.Request) {
	var req ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	p := &model.Product{
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryIDs: req.CategoryIDs,
	}

	ctx := r.Context()
	err := h.service.Create(ctx, p)
	if err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusCreated, true, "product created", nil)
}

func (h *ProductHandler) getByID(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	data, err := h.service.GetByID(ctx, id)
	if err == service.ErrProductNotFound {
		model.JSONResponse(w, http.StatusNotFound, false, "product not found", nil)
		return
	}

	if err != nil {
		model.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusOK, true, "product found", data)
}

func (h *ProductHandler) update(w http.ResponseWriter, r *http.Request, id string) {
	var req ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	p := &model.Product{
		ID:          id,
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryIDs: req.CategoryIDs,
	}

	ctx := r.Context()
	err := h.service.Update(ctx, p)
	if err == service.ErrProductNotFound {
		model.JSONResponse(w, http.StatusNotFound, false, "product not found", nil)
		return
	}

	if err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusOK, true, "product updated", nil)
}

func (h *ProductHandler) delete(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()
	err := h.service.Delete(ctx, id)
	if err == service.ErrProductNotFound {
		model.JSONResponse(w, http.StatusNotFound, false, "product not found", nil)
		return
	}

	if err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusOK, true, "product deleted", nil)
}

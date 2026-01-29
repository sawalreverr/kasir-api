package handler

import (
	"basic-go-api/internal/model"
	"basic-go-api/internal/service"
	"encoding/json"
	"net/http"
	"strings"
)

type CategoryRequest struct {
	Name string `json:"name"`
}

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

func (h *CategoryHandler) Categories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAll(w)
	case http.MethodPost:
		h.create(w, r)
	default:
		model.JSONResponse(w, http.StatusMethodNotAllowed, false, "method not allowed", nil)
	}
}

func (h *CategoryHandler) CategoryByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/categories/")
	if id == "" {
		model.JSONResponse(w, http.StatusNotFound, false, "not found", nil)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getByID(w, id)
	case http.MethodPut:
		h.update(w, r, id)
	case http.MethodDelete:
		h.delete(w, id)
	default:
		model.JSONResponse(w, http.StatusMethodNotAllowed, false, "method not allowed", nil)
	}
}

// GET /categories
func (h *CategoryHandler) getAll(w http.ResponseWriter) {
	data, err := h.service.GetAll()
	if err != nil {
		model.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusOK, true, "success", data)
}

// POST /categories
func (h *CategoryHandler) create(w http.ResponseWriter, r *http.Request) {
	var req CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	err := h.service.Create(req.Name)
	if err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusCreated, true, "category created", nil)
}

// GET /categories/{id}
func (h *CategoryHandler) getByID(w http.ResponseWriter, id string) {
	data, err := h.service.GetByID(id)
	if err == service.ErrCategoryNotFound {
		model.JSONResponse(w, http.StatusNotFound, false, "category not found", nil)
		return
	}

	if err != nil {
		model.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusOK, true, "category found", data)
}

// PUT /categories/{id}
func (h *CategoryHandler) update(w http.ResponseWriter, r *http.Request, id string) {
	var req CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	err := h.service.Update(id, req.Name)
	if err == service.ErrCategoryNotFound {
		model.JSONResponse(w, http.StatusNotFound, false, "category not found", nil)
		return
	}

	if err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusNoContent, true, "category updated", nil)
}

// GET /categories/{id}
func (h *CategoryHandler) delete(w http.ResponseWriter, id string) {
	err := h.service.Delete(id)
	if err == service.ErrCategoryNotFound {
		model.JSONResponse(w, http.StatusNotFound, false, "category not found", nil)
		return
	}

	if err != nil {
		model.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusNoContent, true, "category deleted", nil)
}

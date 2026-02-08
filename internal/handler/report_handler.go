package handler

import (
	"basic-go-api/internal/model"
	"basic-go-api/internal/service"
	"net/http"
	"strings"
)

type ReportHandler struct {
	service *service.ReportService
}

func NewReportHandler(s *service.ReportService) *ReportHandler {
	return &ReportHandler{service: s}
}

func (h *ReportHandler) Reports(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case strings.HasSuffix(path, "/today"):
		h.Today(w, r)
	case strings.HasPrefix(path, "/report"):
		h.DateRange(w, r)
	default:
		model.JSONResponse(w, http.StatusNotFound, false, "not found", nil)
	}
}

func (h *ReportHandler) Today(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		model.JSONResponse(w, http.StatusMethodNotAllowed, false, "method not allowed", nil)
		return
	}

	ctx := r.Context()
	report, err := h.service.GetTodayReport(ctx)
	if err != nil {
		model.JSONResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusOK, true, "today report", report)
}

func (h *ReportHandler) DateRange(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		model.JSONResponse(w, http.StatusMethodNotAllowed, false, "method not allowed", nil)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	ctx := r.Context()
	report, err := h.service.GetDateRangeReport(ctx, startDate, endDate)
	if err != nil {
		model.JSONResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	model.JSONResponse(w, http.StatusOK, true, "report", report)
}

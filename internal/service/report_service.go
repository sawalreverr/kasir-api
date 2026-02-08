package service

import (
	"basic-go-api/internal/model"
	"basic-go-api/internal/repository"
	"context"
	"errors"
	"time"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService(repo *repository.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodayReport(ctx context.Context) (*model.Report, error) {
	now := time.Now()

	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 0, 1)

	return s.repo.GetReport(ctx, start, end)
}

func (s *ReportService) GetDateRangeReport(ctx context.Context, startDate, endDate string) (*model.Report, error) {
	if startDate == "" || endDate == "" {
		return nil, errors.New("start_date and end_date are required")
	}

	start, err := time.ParseInLocation("2006-01-02", startDate, time.Local)
	if err != nil {
		return nil, errors.New("invalid start_date format, use YYYY-MM-DD")
	}

	end, err := time.ParseInLocation("2006-01-02", endDate, time.Local)
	if err != nil {
		return nil, errors.New("invalid end_date format, use YYYY-MM-DD")
	}

	if start.After(end) {
		return nil, errors.New("start_date cannot be after end_date")
	}

	end = end.AddDate(0, 0, 1)

	return s.repo.GetReport(ctx, start, end)
}

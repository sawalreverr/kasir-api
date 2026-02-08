package repository

import (
	"basic-go-api/internal/model"
	"context"
	"database/sql"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetReport(ctx context.Context, startDate, endDate time.Time) (*model.Report, error) {
	var totalRevenue int
	var totalTransaction int

	err := r.db.QueryRowContext(ctx, `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`, startDate, endDate).Scan(&totalRevenue, &totalTransaction)
	if err != nil {
		return nil, err
	}

	bestSeller := &model.BestSeller{}
	row := r.db.QueryRowContext(ctx, `
		SELECT p.name, SUM(td.quantity)
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		JOIN products p ON p.id = td.product_id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.id, p.name
		ORDER BY SUM(td.quantity) DESC
		LIMIT 1
	`, startDate, endDate)

	err = row.Scan(&bestSeller.Name, &bestSeller.QtySold)
	if err != nil {
		if err == sql.ErrNoRows {
			bestSeller = nil
		} else {
			return nil, err
		}
	}

	return &model.Report{
		TotalRevenue:     totalRevenue,
		TotalTransaction: totalTransaction,
		BestSeller:       bestSeller,
	}, nil
}

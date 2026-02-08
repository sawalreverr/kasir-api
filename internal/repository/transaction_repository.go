package repository

import (
	"basic-go-api/internal/model"
	"context"
	"database/sql"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, tx *sql.Tx, t *model.Transaction) error {
	err := tx.QueryRowContext(ctx, `
		INSERT INTO transactions (total_amount)
		VALUES ($1)
		RETURNING id, created_at
	`, t.TotalAmount).Scan(&t.ID, &t.CreatedAt)
	if err != nil {
		return err
	}

	for i := range t.Details {
		detail := &t.Details[i]
		detail.TransactionID = t.ID
		err = tx.QueryRowContext(ctx, `
			INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`, detail.TransactionID, detail.ProductID, detail.Quantity, detail.Subtotal).Scan(&detail.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

package service

import (
	"basic-go-api/internal/model"
	"basic-go-api/internal/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type TransactionService struct {
	repo        *repository.TransactionRepository
	productRepo *repository.ProductRepository
	db          *sql.DB
}

func NewTransactionService(repo *repository.TransactionRepository, productRepo *repository.ProductRepository, db *sql.DB) *TransactionService {
	return &TransactionService{
		repo:        repo,
		productRepo: productRepo,
		db:          db,
	}
}

func (s *TransactionService) Create(ctx context.Context, items []model.CheckoutItem) (*model.Transaction, error) {
	if len(items) == 0 {
		return nil, errors.New("transaction must have at least one item")
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]model.TransactionDetail, 0, len(items))

	for _, item := range items {
		if item.ProductID <= 0 {
			return nil, errors.New("product_id must be greater than 0")
		}
		if item.Quantity <= 0 {
			return nil, errors.New("quantity must be greater than 0")
		}

		product, err := s.productRepo.FindByID(ctx, fmt.Sprintf("%d", item.ProductID))
		if err != nil {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if product == nil {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}

		err = s.productRepo.UpdateStock(ctx, tx, product.ID, item.Quantity)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("insufficient stock for product %s", product.Name)
		}
		if err != nil {
			return nil, err
		}

		subtotal := product.Price * item.Quantity
		totalAmount += subtotal

		details = append(details, model.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	transaction := &model.Transaction{
		TotalAmount: totalAmount,
		Details:     details,
	}

	err = s.repo.CreateTransaction(ctx, tx, transaction)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return transaction, nil
}

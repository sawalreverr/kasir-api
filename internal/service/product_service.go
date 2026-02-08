package service

import (
	"basic-go-api/internal/model"
	"basic-go-api/internal/repository"
	"context"
	"database/sql"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(r *repository.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*model.Product, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, ErrProductNotFound
	}

	return p, nil
}

func (s *ProductService) GetAll(ctx context.Context, name string) ([]model.Product, error) {
	return s.repo.FindAll(ctx, name)
}

func (s *ProductService) Create(ctx context.Context, p *model.Product) error {
	if p.Name == "" {
		return ErrProductNameEmpty
	}

	if p.Price <= 0 {
		return ErrProductPriceInvalid
	}

	return s.repo.Create(ctx, p)
}

func (s *ProductService) Update(ctx context.Context, p *model.Product) error {
	if p.Name == "" {
		return ErrProductNameEmpty
	}

	if p.Price <= 0 {
		return ErrProductPriceInvalid
	}

	err := s.repo.Update(ctx, p)
	if err == sql.ErrNoRows {
		return ErrProductNotFound
	}

	return err
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err == sql.ErrNoRows {
		return ErrProductNotFound
	}

	return err
}

package service

import (
	"basic-go-api/internal/model"
	"basic-go-api/internal/repository"
	"database/sql"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(r *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: r}
}

func (s *CategoryService) GetByID(id string) (*model.Category, error) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, ErrCategoryNotFound
	}

	return c, nil
}

func (s *CategoryService) Create(name, description string) error {
	if name == "" {
		return ErrCategoryNameEmpty
	}

	return s.repo.Create(name, description)
}

func (s *CategoryService) Update(id, name, description string) error {
	if name == "" {
		return ErrCategoryNameEmpty
	}

	err := s.repo.Update(id, name, description)
	if err == sql.ErrNoRows {
		return ErrCategoryNotFound
	}
	return err
}

func (s *CategoryService) Delete(id string) error {
	err := s.repo.Delete(id)
	if err == sql.ErrNoRows {
		return ErrCategoryNotFound
	}
	return err
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repo.FindAll()
}

package repository

import (
	"basic-go-api/internal/model"
	"database/sql"
	"errors"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindByID(id string) (*model.Category, error) {
	var c model.Category
	err := r.db.QueryRow(`SELECT id, name FROM categories WHERE id = $1 AND deleted_at IS NULL`, id).Scan(&c.ID, &c.Name)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &c, err
}

func (r *CategoryRepository) Create(name string) error {
	_, err := r.db.Exec(`INSERT INTO categories (name) VALUES ($1)`, name)
	return err
}

func (r *CategoryRepository) Update(id, name string) error {
	result, err := r.db.Exec(`UPDATE categories SET name = $1, updated_at = now() WHERE id = $2 AND deleted_at IS NULL`, name, id)
	if err != nil {
		return nil
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *CategoryRepository) Delete(id string) error {
	result, err := r.db.Exec(`UPDATE categories SET deleted_at = now() WHERE id = $1 AND deleted_at IS NULL`, id)
	if err != nil {
		return nil
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	rows, err := r.db.Query(`SELECT id, name FROM categories WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		rows.Scan(&c.ID, &c.Name)
		categories = append(categories, c)
	}

	return categories, nil
}

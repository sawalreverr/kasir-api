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
	err := r.db.QueryRow(`SELECT id, name, description FROM categories WHERE id = $1 AND deleted_at IS NULL`, id).Scan(&c.ID, &c.Name, &c.Description)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &c, err
}

func (r *CategoryRepository) Create(name, description string) error {
	_, err := r.db.Exec(`INSERT INTO categories (name, description) VALUES ($1, $2)`, name, description)
	return err
}

func (r *CategoryRepository) Update(id, name, description string) error {
	result, err := r.db.Exec(`UPDATE categories SET name = $1, description = $2, updated_at = now() WHERE id = $3 AND deleted_at IS NULL`, name, description, id)
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
	rows, err := r.db.Query(`SELECT id, name, description FROM categories WHERE deleted_at IS NULL ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var c model.Category
		rows.Scan(&c.ID, &c.Name, &c.Description)
		categories = append(categories, c)
	}

	return categories, nil
}

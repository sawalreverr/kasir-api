package repository

import (
	"basic-go-api/internal/model"
	"context"
	"database/sql"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindByID(ctx context.Context, id string) (*model.Product, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT
			p.id, p.name, p.price, p.stock,
			c.id, c.name, c.description
		FROM products p
		LEFT JOIN product_categories pc ON pc.product_id = p.id
		LEFT JOIN categories c ON c.id = pc.category_id AND c.deleted_at IS NULL
		WHERE p.id = $1 AND p.deleted_at IS NULL
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product *model.Product
	for rows.Next() {
		var (
			pID, pName        string
			price, stock      int
			cID, cName, cDesc sql.NullString
		)

		err := rows.Scan(
			&pID, &pName, &price, &stock,
			&cID, &cName, &cDesc,
		)
		if err != nil {
			return nil, err
		}

		if product == nil {
			product = &model.Product{
				ID:    pID,
				Name:  pName,
				Price: price,
				Stock: stock,
			}
		}

		if cID.Valid {
			product.Categories = append(product.Categories, model.Category{
				ID:          cID.String,
				Name:        cName.String,
				Description: cDesc.String,
			})
		}
	}

	if product == nil {
		return nil, nil
	}

	return product, nil
}

func (r *ProductRepository) Create(ctx context.Context, p *model.Product) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, `
		INSERT INTO products (name, price, stock)
		VALUES ($1, $2, $3)
		RETURNING id
	`, p.Name, p.Price, p.Stock).Scan(&p.ID)
	if err != nil {
		return err
	}

	for _, categoryID := range p.CategoryIDs {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO product_categories (product_id, category_id)
			VALUES ($1, $2)
		`, p.ID, categoryID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *ProductRepository) Update(ctx context.Context, p *model.Product) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, `
		UPDATE products
		SET name = $1, price = $2, stock = $3, updated_at = now()
		WHERE id = $4 AND deleted_at IS NULL
	`, p.Name, p.Price, p.Stock, p.ID)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}

	_, err = tx.ExecContext(ctx, `
		DELETE FROM product_categories WHERE product_id = $1
	`, p.ID)
	if err != nil {
		return err
	}

	for _, categoryID := range p.CategoryIDs {
		_, err = tx.ExecContext(ctx, `
			INSERT INTO product_categories (product_id, category_id)
			VALUES ($1, $2)
		`, p.ID, categoryID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	result, err := r.db.ExecContext(ctx, `
		UPDATE products SET deleted_at = now()
		WHERE id = $1 AND deleted_at IS NULL
	`, id)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT
			p.id, p.name, p.price, p.stock,
			pc.category_id
		FROM products p
		LEFT JOIN product_categories pc ON pc.product_id = p.id
		WHERE p.deleted_at IS NULL
		ORDER BY p.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productMap := make(map[string]*model.Product)
	var products []*model.Product

	for rows.Next() {
		var (
			pID, name    string
			price, stock int
			categoryID   sql.NullString
		)

		err := rows.Scan(&pID, &name, &price, &stock, &categoryID)
		if err != nil {
			return nil, err
		}

		product, exists := productMap[pID]
		if !exists {
			product = &model.Product{
				ID:    pID,
				Name:  name,
				Price: price,
				Stock: stock,
			}
			productMap[pID] = product
			products = append(products, product)
		}

		if categoryID.Valid {
			product.CategoryIDs = append(product.CategoryIDs, categoryID.String)
		}
	}

	var result []model.Product
	for _, p := range products {
		result = append(result, *p)
	}

	return result, nil
}

package seed

import (
	"context"
	"database/sql"
	"log"
)

func Run(db *sql.DB) {
	ctx := context.Background()

	log.Println("[seed] seeding categories...")
	_, _ = db.ExecContext(ctx, `
		INSERT INTO categories (name, description)
		VALUES
		  ('Electronics', 'Electronic devices'),
		  ('Books', 'All kinds of books'),
		  ('Clothing', 'Wearable items')
		ON CONFLICT DO NOTHING
	`)

	log.Println("[seed] seeding products...")
	_, _ = db.ExecContext(ctx, `
		INSERT INTO products (name, price, stock)
		VALUES
		  ('Laptop', 15000000, 10),
		  ('T-Shirt', 150000, 50)
		ON CONFLICT DO NOTHING
	`)

	log.Println("[seed] seeding product_categories...")
	_, _ = db.ExecContext(ctx, `
		INSERT INTO product_categories (product_id, category_id)
		VALUES
		  ('1', '1'),
		  ('2', '3')
		ON CONFLICT DO NOTHING
	`)

	log.Println("[seed] done")
}

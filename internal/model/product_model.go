package model

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`

	CategoryIDs []string   `json:"category_ids,omitempty"`
	Categories  []Category `json:"categories,omitempty"`
}

package model

type Report struct {
	TotalRevenue     int         `json:"total_revenue"`
	TotalTransaction int         `json:"total_transaction"`
	BestSeller       *BestSeller `json:"best_seller"`
}

type BestSeller struct {
	Name    string `json:"name"`
	QtySold int    `json:"qty_sold"`
}

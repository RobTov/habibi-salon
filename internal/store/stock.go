package store

import "database/sql"

type Stock struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	ServiceID int64 `json:"service_id"`
	Quantity  int64 `json:"quantity"`
}

type StockStore struct {
	db *sql.DB
}

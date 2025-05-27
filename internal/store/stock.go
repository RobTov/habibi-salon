package store

import "database/sql"

type Stock struct {
}

type StockStore struct {
	db *sql.DB
}

package store

import (
	"database/sql"
)

type Services struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	IsActive    string  `json:"is_active"`
}

type ServicesStore struct {
	db *sql.DB
}

package store

import (
	"database/sql"
	"time"
)

type Clients struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type ClientsStore struct {
	db *sql.DB
}

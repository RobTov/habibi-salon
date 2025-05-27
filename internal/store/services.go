package store

import "database/sql"

type Services struct {
}

type ServicesStore struct {
	db *sql.DB
}

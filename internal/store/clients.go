package store

import "database/sql"

type Clients struct {
}

type ClientsStore struct {
	db *sql.DB
}

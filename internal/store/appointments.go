package store

import "database/sql"

type Appointments struct {
}

type AppointmentsStore struct {
	db *sql.DB
}

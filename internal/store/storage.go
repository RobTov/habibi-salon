package store

import (
	"database/sql"
	"time"
)

var (
	// ErrNotFound          = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Appointments interface{}
	Clients      interface{}
	Services     interface{}
	Stock        interface{}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Appointments: &AppointmentsStore{db},
		Clients:      &ClientsStore{db},
		Services:     &ServicesStore{db},
		Stock:        &StockStore{db},
	}
}

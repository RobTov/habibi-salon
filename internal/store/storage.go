package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Appointments interface {
		Create(context.Context, *Appointments) error
		GetAll(context.Context) ([]Appointments, error)
		GetByID(context.Context, int64) (*Appointments, error)
		Delete(context.Context, int64) error
	}
	Clients interface {
		Create(context.Context, *Clients) error
		GetByID(context.Context, int64) (*Clients, error)
		Update(context.Context, *Clients) error
		Delete(context.Context, int64) error
	}
	Services interface{}
	Stock    interface{}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Appointments: &AppointmentsStore{db},
		Clients:      &ClientsStore{db},
		Services:     &ServicesStore{db},
		Stock:        &StockStore{db},
	}
}

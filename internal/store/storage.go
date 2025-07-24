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
		GetAll(context.Context) ([]Clients, error)
		Create(context.Context, *Clients) error
		GetByID(context.Context, int64) (*Clients, error)
		Update(context.Context, *Clients) error
		Delete(context.Context, int64) error
	}
	Services interface {
		GetAll(context.Context) ([]Services, error)
		GetByID(context.Context, int64) (*Services, error)
		Update(context.Context, *Services) error
		Create(context.Context, *Services) error
		Delete(context.Context, int64) error
	}
	Products interface {
		GetAll(context.Context) ([]Products, error)
		Create(context.Context, *Products) error
	}
	Stock interface {
		GetAll(context.Context) ([]Stock, error)
		Create(context.Context, *Stock) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Appointments: &AppointmentsStore{db},
		Clients:      &ClientsStore{db},
		Services:     &ServicesStore{db},
		Products:     &ProductsStore{db},
		Stock:        &StockStore{db},
	}
}

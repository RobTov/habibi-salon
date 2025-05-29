package store

import (
	"context"
	"database/sql"
	"errors"
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

func (s *ClientsStore) Create(ctx context.Context, client *Clients) error {
	query := `
	INSERT INTO clients (name, email, phone, address) 
	VALUES ($1, $2, $3, $4);
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, client.Name, client.Phone, client.Address)
	if err != nil {
		return err
	}

	return nil
}

func (s *ClientsStore) GetByID(ctx context.Context, clientID int64) (*Clients, error) {
	query := `SELECT id, name, email, phone, address, created_at 
	FROM clients WHERE id = $1;`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var client Clients
	err := s.db.QueryRowContext(ctx, query, clientID).Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.Phone,
		&client.Address,
		&client.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &client, nil
}

func (s *ClientsStore) Update(ctx context.Context, client *Clients) error {
	query := `
	UPDATE clients 
	SET name = $1, email = $2, phone = $3, address = $4
	WHERE id = $5;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(
		ctx,
		query,
		client.Name,
		client.Email,
		client.Phone,
		client.Address,
		client.ID,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *ClientsStore) Delete(ctx context.Context, clientID int64) error {
	query := `
	DELETE FROM clients WHERE id = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, clientID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

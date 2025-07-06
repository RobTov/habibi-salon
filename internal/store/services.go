package store

import (
	"context"
	"database/sql"
)

type Services struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	IsActive    bool    `json:"is_active"`
}

type ServicesStore struct {
	db *sql.DB
}

func (s *ServicesStore) GetAll(ctx context.Context) ([]Services, error) {
	query := `
	SELECT id, name, description, price, is_active
	FROM services;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	services := []Services{}
	for rows.Next() {
		var s = Services{}

		err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Description,
			&s.Price,
			&s.IsActive,
		)

		if err != nil {
			return nil, err
		}

		services = append(services, s)

	}

	return services, nil
}

func (s *ServicesStore) Create(ctx context.Context, service *Services) error {
	query := `
	INSERT INTO services (name, description, price)
	VALUES ($1, $2, $3);
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, service.Name, service.Description, service.Price)

	if err != nil {
		return err
	}

	return nil
}

func (s *ServicesStore) Delete(ctx context.Context, serviceID int64) error {
	query := `
	DELETE FROM services WHERE id = $1;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, serviceID)

	if err != nil {
		return err
	}

	return nil
}

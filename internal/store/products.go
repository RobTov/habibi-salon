package store

import (
	"context"
	"database/sql"
)

type Products struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
}

type ProductsStore struct {
	db *sql.DB
}

func (s ProductsStore) GetAll(ctx context.Context) ([]Products, error) {
	query := `
	SELECT id, name, description, quantity FROM products;
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Products{}
	for rows.Next() {
		var p = Products{}

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Quantity,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
